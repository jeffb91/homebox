package v1

import (
	"errors"
	"net/http"

	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog/log"
	"github.com/sysadminsmedia/homebox/backend/internal/core/services"
	"github.com/sysadminsmedia/homebox/backend/internal/data/repo"
	"github.com/sysadminsmedia/homebox/backend/internal/sys/validate"
)

// POST /v1/maintenance/{id}/attachments
func (ctrl *V1Controller) HandleMaintenanceAttachmentCreate() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseMultipartForm(ctrl.maxUploadSize << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			return validate.NewRequestError(errors.New("failed to parse form"), http.StatusBadRequest)
		}

		errs := validate.NewFieldErrors()

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			errs = errs.Append("file", "file is required")
		}

		filename := r.FormValue("name")
		if filename == "" {
			errs = errs.Append("name", "name is required")
		}

		if !errs.Nil() {
			return server.JSON(w, http.StatusUnprocessableEntity, errs)
		}
		defer file.Close()

		entryID, err := ctrl.routeUUID(r, "id")
		if err != nil {
			return err
		}

		ctx := services.NewContext(r.Context())

		_, err = ctrl.repo.MaintAttachments.Create(ctx, entryID, repo.MaintenanceAttachmentCreate{
			Filename:    filename,
			ContentType: fileHeader.Header.Get("Content-Type"),
			Content:     file,
		})
		if err != nil {
			log.Err(err).Msg("failed to save maintenance attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

// GET /v1/maintenance/attachments/{attachment_id}
func (ctrl *V1Controller) HandleMaintenanceAttachmentGet() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return err
		}

		ctx := services.NewContext(r.Context())

		attachment, err := ctrl.repo.MaintAttachments.Get(ctx, attachmentID)
		if err != nil {
			return validate.NewRequestError(err, http.StatusNotFound)
		}

		http.ServeFile(w, r, attachment.Path)
		return nil
	}
}

// GET /v1/maintenance/{id}/attachments
func (ctrl *V1Controller) HandleMaintenanceAttachmentList() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		entryID, err := ctrl.routeUUID(r, "id")
		if err != nil {
			return err
		}

		ctx := services.NewContext(r.Context())

		attachments, err := ctrl.repo.MaintAttachments.ListByEntry(ctx, entryID)
		if err != nil {
			log.Err(err).Msg("failed to list maintenance attachments")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, attachments)
	}
}

// DELETE /v1/maintenance/attachments/{attachment_id}
func (ctrl *V1Controller) HandleMaintenanceAttachmentDelete() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return err
		}

		ctx := services.NewContext(r.Context())

		err = ctrl.repo.MaintAttachments.Delete(ctx, attachmentID)
		if err != nil {
			log.Err(err).Msg("failed to delete maintenance attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, map[string]string{"status": "deleted"})
	}
}

// PUT /v1/maintenance/attachments/{attachment_id}
func (ctrl *V1Controller) HandleMaintenanceAttachmentUpdate() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		attachmentID, err := ctrl.routeUUID(r, "attachment_id")
		if err != nil {
			return err
		}

		var data struct {
			Title string `json:"title"`
		}

		if err := server.DecodeJSON(r.Body, &data); err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		ctx := services.NewContext(r.Context())

		_, err = ctrl.repo.MaintAttachments.Rename(ctx, attachmentID, data.Title)
		if err != nil {
			log.Err(err).Msg("failed to update attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, map[string]string{"status": "updated"})
	}
}
