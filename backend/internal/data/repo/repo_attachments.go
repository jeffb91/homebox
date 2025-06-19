package repo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/attachment"
)

type AttachmentRepo struct {
	db  *ent.Client
	dir string
}

type (
	Attachment struct {
		ID          uuid.UUID `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Type        string    `json:"type"`
		Primary     bool      `json:"primary"`
		Path        string    `json:"path"`
		Title       string    `json:"title"`
		RelatedType string    `json:"related_type"`
		RelatedID   uuid.UUID `json:"related_id"`
	}

	AttachmentCreate struct {
		RelatedType string
		RelatedID   uuid.UUID
		Title       string
		Type        string
		File        io.Reader
		Primary     bool
	}

	AttachmentUpdate struct {
		Title   string
		Type    string
		Primary bool
	}
)

/* AttachmentRepo struct definition removed to avoid redeclaration error.
   Please ensure AttachmentRepo is defined only once, e.g., in repo_item_attachments.go. */

// Helper om het pad te bepalen waar het bestand wordt opgeslagen
//
//	func (r *AttachmentRepo) path(relatedType string, relatedID uuid.UUID, filename string) string {
//		return filepath.Join(r.dir, relatedType, relatedID.String(), "documents", filename)
//	}
func (r *AttachmentRepo) path(relatedType string, relatedID uuid.UUID, filename string) string {
	baseDir := filepath.Join(r.dir, "attachments", fmt.Sprintf("item-%s", relatedID.String()))
	var subDir string

	switch relatedType {
	case "documents":
		subDir = "documents"
	case "maintenance":
		subDir = "maintenance"
	case "issues":
		subDir = "issues"
	default:
		subDir = "documents"
	}

	// Bestandsnaam prefixen met type en id voor uniekheid
	prefixedFilename := fmt.Sprintf("%s-%s", relatedType, filename)

	return filepath.Join(baseDir, subDir, prefixedFilename)
}

// Create een attachment en sla het bestand op
func (r *AttachmentRepo) Create(ctx context.Context, input AttachmentCreate) (*ent.Attachment, error) {
	// Lees de file in een buffer
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, input.File)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Genereer een unieke bestandsnaam (bijv. uuid + extensie)
	filename := fmt.Sprintf("%s_%s", uuid.New().String(), input.Title)
	filePath := r.path(input.RelatedType, input.RelatedID, filename)
	//filePath := r.path("maintenance", itemID, "remcontrole.pdf")

	// Maak de directory aan als die nog niet bestaat
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Sla het bestand op
	if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	// Maak het attachment record aan in de database
	att, err := r.db.Attachment.
		Create().
		SetRelatedType(input.RelatedType).
		SetRelatedID(input.RelatedID).
		SetTitle(input.Title).
		SetType(attachment.Type(input.Type)).
		SetPrimary(input.Primary).
		SetPath(filePath).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create attachment record: %w", err)
	}

	return att, nil
}

func (r *AttachmentRepo) Get(ctx context.Context, id uuid.UUID) (*ent.Attachment, error) {
	return r.db.Attachment.Query().Where(attachment.ID(id)).Only(ctx)
}

func (r *AttachmentRepo) Update(ctx context.Context, id uuid.UUID, input AttachmentUpdate) (*ent.Attachment, error) {
	upd := r.db.Attachment.UpdateOneID(id)
	if input.Title != "" {
		upd.SetTitle(input.Title)
	}
	if input.Type != "" {
		upd.SetType(attachment.Type(input.Type))
	}
	upd.SetPrimary(input.Primary)
	att, err := upd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update attachment: %w", err)
	}
	return att, nil
}

func (r *AttachmentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	// Haal het attachment op om het bestandspad te weten
	att, err := r.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get attachment: %w", err)
	}

	// Verwijder het bestand
	if att.Path != "" {
		_ = os.Remove(att.Path) // negeer fout als bestand niet bestaat
	}

	// Verwijder het record uit de database
	return r.db.Attachment.DeleteOneID(id).Exec(ctx)
}

func (r *AttachmentRepo) ListByRelated(ctx context.Context, relatedType string, relatedID uuid.UUID) ([]*ent.Attachment, error) {
	return r.db.Attachment.Query().
		Where(
			attachment.RelatedTypeEQ(relatedType),
			attachment.RelatedID(relatedID),
		).
		All(ctx)
}
