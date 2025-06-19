package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent"
	"github.com/sysadminsmedia/homebox/backend/internal/data/repo"
)

type AttachmentService struct {
	repo *repo.AllRepos
}

func (svc *AttachmentService) AttachmentPath(ctx context.Context, attachmentID uuid.UUID) (*ent.Attachment, error) {
	attachment, err := svc.repo.Attachments.Get(ctx, attachmentID)
	if err != nil {
		return nil, err
	}

	return attachment, nil
}

func (svc *AttachmentService) AttachmentUpdate(ctx context.Context, attachmentID uuid.UUID, data repo.AttachmentUpdate) (*ent.Attachment, error) {
	return svc.repo.Attachments.Update(ctx, attachmentID, data)
}

//func (svc *AttachmentService) AttachmentUpdate(ctx Context, itemID, attachmentID uuid.UUID, data repo.AttachmentUpdate) (repo.ItemOut, error) {
//	// Update Attachment
//	attachment, err := svc.repo.Attachments.Update(ctx, attachmentID, data)
//	if err != nil {
//		return repo.ItemOut{}, err
//	}
//
//	// Update Document
//	attDoc := attachment
//	_, err = svc.repo.Attachments.Rename(ctx, attDoc.ID, data.Title)
//	if err != nil {
//		return repo.ItemOut{}, err
//	}
//
//	return svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
//}

// AttachmentAdd adds an attachment to an item by creating an entry in the Documents table and linking it to the Attachment
// Table and Items table. The file provided via the reader is stored on the file system based on the provided
// relative path during construction of the service.
//func (svc *AttachmentService) AttachmentAdd(ctx Context, itemID uuid.UUID, filename string, attachmentType attachment.Type, primary bool, file io.Reader) (repo.ItemOut, error) {
//	// Get the Item
//	_, err := svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
//	if err != nil {
//		return repo.ItemOut{}, err
//	}
//
//	// Create the attachment
//	//_, err = svc.repo.Attachments.Create(ctx, itemID, repo.AttachmentCreate{Title: filename, Content: file}, attachmentType, primary)
//
//	_, err = svc.repo.Attachments.Create(ctx, repo.AttachmentCreate{
//		RelatedType: "item",
//		RelatedID:   itemID,
//		Title:       filename,
//		Type:        string(attachmentType),
//		File:        file,
//		Primary:     primary,
//	})
//
//	if err != nil {
//		log.Err(err).Msg("failed to create attachment")
//	}
//
//	return svc.repo.Items.GetOneByGroup(ctx, ctx.GID, itemID)
//}

func (svc *AttachmentService) AttachmentAdd(ctx context.Context, input repo.AttachmentCreate) (*ent.Attachment, error) {
	return svc.repo.Attachments.Create(ctx, input)
}

//func (svc *AttachmentService) AttachmentDelete(ctx context.Context, gid, itemID, attachmentID uuid.UUID) error {
//	// Delete the attachment
//	err := svc.repo.Attachments.Delete(ctx, attachmentID)
//	if err != nil {
//		return err
//	}
//
//	return err
//}

func (svc *AttachmentService) AttachmentDelete(ctx context.Context, attachmentID uuid.UUID) error {
	return svc.repo.Attachments.Delete(ctx, attachmentID)
}
