package v1

import (
	"errors"
	"net/http"
	"path/filepath"

	//	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog/log"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/attachment"
	"github.com/sysadminsmedia/homebox/backend/internal/data/repo"
	"github.com/sysadminsmedia/homebox/backend/internal/sys/validate"
)

type (
	AttachmentToken struct {
		Token string `json:"token"`
	}
)

// HandleAttachmentCreate godocs
//
//	@Summary  Create Attachment
//	@Tags     Attachments
//	@Accept   multipart/form-data
//	@Produce  json
//	@Param    related_type formData string true "Type of related entity (e.g. item, maintenance)"
//	@Param    related_id   formData string true "ID of the related entity"
//	@Param    file         formData file   true "File attachment"
//	@Param    type         formData string true "Type of file"
//	@Param    name         formData string true "Name of the file including extension"
//	@Success  201  {object} repo.Attachment
//	@Failure  422  {object} validate.ErrorResponse
//	@Router   /v1/attachments [POST]
//	@Security Bearer
func (ctrl *V1Controller) HandleAttachmentCreate() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseMultipartForm(ctrl.maxUploadSize << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			return validate.NewRequestError(errors.New("failed to parse multipart form"), http.StatusBadRequest)
		}

		errs := validate.NewFieldErrors()

		file, _, err := r.FormFile("file")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrMissingFile):
				log.Debug().Msg("file for attachment is missing")
				errs = errs.Append("file", "file is required")
			default:
				log.Err(err).Msg("failed to get file from form")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
		}

		attachmentName := r.FormValue("name")
		if attachmentName == "" {
			log.Debug().Msg("failed to get name from form")
			errs = errs.Append("name", "name is required")
		}

		relatedType := r.FormValue("related_type")
		relatedID := r.FormValue("related_id")
		var relatedUUID uuid.UUID
		if relatedType == "" {
			errs = errs.Append("related_type", "related_type is required")
		}
		if relatedID == "" {
			errs = errs.Append("related_id", "related_id is required")
		} else {
			var err error
			relatedUUID, err = uuid.Parse(relatedID)
			if err != nil {
				errs = errs.Append("related_id", "invalid related_id")
			}
		}

		if !errs.Nil() {
			return server.JSON(w, http.StatusUnprocessableEntity, errs)
		}

		attachmentType := r.FormValue("type")
		if attachmentType == "" {
			// Attempt to auto-detect the type of the file
			ext := filepath.Ext(attachmentName)

			switch strings.ToLower(ext) {
			case ".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp", ".tiff":
				attachmentType = attachment.TypePhoto.String()
			default:
				attachmentType = attachment.TypeAttachment.String()
			}
		}

		// Optionally handle "primary" field if needed in the future

		_, err = ctrl.routeID(r)
		if err != nil {
			return err
		}

		attachment, err := ctrl.repo.Attachments.Create(r.Context(), repo.AttachmentCreate{
			RelatedType: relatedType,
			RelatedID:   relatedUUID,
			Title:       attachmentName,
			Type:        attachmentType,
			File:        file,
			Primary:     false,
		})
		if err != nil {
			log.Err(err).Msg("failed to add attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}
		return server.JSON(w, http.StatusCreated, attachment)
	}
}

// GET handler
func (ctrl *V1Controller) HandleAttachmentGet() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		att, err := ctrl.repo.Attachments.Get(r.Context(), attachmentID)
		if err != nil {
			log.Err(err).Msg("failed to get attachment")
			return validate.NewRequestError(err, http.StatusNotFound)
		}

		http.ServeFile(w, r, att.Path)
		return nil
	}
}

// DELETE handler
func (ctrl *V1Controller) HandleAttachmentDelete() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		err = ctrl.repo.Attachments.Delete(r.Context(), attachmentID)
		if err != nil {
			log.Err(err).Msg("failed to delete attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

// UPDATE handler
func (ctrl *V1Controller) HandleAttachmentUpdate() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		var update repo.AttachmentUpdate
		err = server.Decode(r, &update)
		if err != nil {
			log.Err(err).Msg("failed to decode attachment update")
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		att, err := ctrl.repo.Attachments.Update(r.Context(), attachmentID, update)
		if err != nil {
			log.Err(err).Msg("failed to update attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, att)
	}
}

// HandleAttachmentList godoc
//
//	@Summary	List Attachments
//	@Tags		Attachments
//	@Produce	json
//	@Param		related_type	query	string	false	"Type of related entity"
//	@Param		related_id		query	string	false	"ID of related entity (UUID)"
//	@Success	200	{object}	[]repo.Attachment
//	@Router		/v1/attachments [GET]
//	@Security	Bearer
func (ctrl *V1Controller) HandleAttachmentList() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		relatedType := r.URL.Query().Get("related_type")
		relatedIDStr := r.URL.Query().Get("related_id")

		var relatedID uuid.UUID
		var err error
		if relatedIDStr != "" {
			relatedID, err = uuid.Parse(relatedIDStr)
			if err != nil {
				return validate.NewRequestError(errors.New("invalid related_id"), http.StatusBadRequest)
			}
		}

		attachments, err := ctrl.repo.Attachments.ListByRelated(r.Context(), relatedType, relatedID)
		if err != nil {
			log.Err(err).Msg("failed to list attachments")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, WrapResults(attachments))
	}
}
