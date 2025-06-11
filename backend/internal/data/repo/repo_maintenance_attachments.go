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
	"github.com/zeebo/blake3"

	"github.com/sysadminsmedia/homebox/backend/internal/data/ent"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/maintenanceattachment"
)

// MaintenanceAttachmentRepo beheert bestandsopslag en DB-opslag voor onderhoudsbijlagen.
type MaintenanceAttachmentRepo struct {
	db  *ent.Client
	dir string // Bijvoorbeeld "data/documents"
}

// Constructor
func NewMaintenanceAttachmentRepo(db *ent.Client, dir string) *MaintenanceAttachmentRepo {
	return &MaintenanceAttachmentRepo{db: db, dir: dir}
}

type MaintenanceAttachmentCreate struct {
	Filename    string
	Content     io.Reader
	ContentType string
}

// maakt het pad voor de opslag van het bestand
func (r *MaintenanceAttachmentRepo) path(entryID uuid.UUID, hash string) string {
	return filepath.Join(r.dir, "maintenance", entryID.String(), hash)
}

func (r *MaintenanceAttachmentRepo) Create(ctx context.Context, entryID uuid.UUID, input MaintenanceAttachmentCreate) (*ent.MaintenanceAttachment, error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
		}
	}()

	// Lees bestand in buffer voor hashing + opslag
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, input.Content)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	contentBytes := buf.Bytes()

	// Genereer hash met maintenance entry ID als context
	hashOut := make([]byte, 32)
	blake3.DeriveKey(entryID.String(), contentBytes, hashOut)
	hashStr := fmt.Sprintf("%x", hashOut)

	// Bepaal pad en maak map aan
	path := r.path(entryID, hashStr)
	parent := filepath.Dir(path)
	if err := os.MkdirAll(parent, 0755); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// Schrijf bestand (indien nog niet aanwezig)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		defer file.Close()

		if _, err = file.Write(contentBytes); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}

	// Sla in database op
	record, err := tx.MaintenanceAttachment.Create().
		SetID(uuid.New()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetFilename(input.Filename).
		SetContentType(input.ContentType).
		SetPath(path).
		SetMaintenanceID(entryID).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *MaintenanceAttachmentRepo) Get(ctx context.Context, id uuid.UUID) (*ent.MaintenanceAttachment, error) {
	return r.db.MaintenanceAttachment.
		Query().
		Where(maintenanceattachment.ID(id)).
		Only(ctx)
}

func (r *MaintenanceAttachmentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	doc, err := r.db.MaintenanceAttachment.Get(ctx, id)
	if err != nil {
		return err
	}

	// Als dit de enige bijlage voor dit pad is: verwijder bestand
	all, err := r.db.MaintenanceAttachment.Query().Where(maintenanceattachment.Path(doc.Path)).All(ctx)
	if err != nil {
		return err
	}
	if len(all) == 1 {
		if err := os.Remove(doc.Path); err != nil {
			return err
		}
	}

	return r.db.MaintenanceAttachment.DeleteOneID(id).Exec(ctx)
}
