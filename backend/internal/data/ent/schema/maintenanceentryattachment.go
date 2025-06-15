package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MaintenanceEntryAttachment holds the schema definition for the attachment.
type MaintenanceEntryAttachment struct {
	ent.Schema
}

// Fields of the MaintenanceEntryAttachment.
func (MaintenanceEntryAttachment) Fields() []ent.Field {
	return []ent.Field{
		// ✅ Primary key als UUID
		field.UUID("id", uuid.New()).Default(uuid.New),

		// ✅ Bestandsgegevens
		field.String("filename").NotEmpty(),
		field.String("filepath").NotEmpty(),
		field.String("content_type").Optional(),

		// ✅ Timestamps
		field.Time("uploaded_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),

		// ✅ Foreign key naar MaintenanceEntry
		field.UUID("maintenance_entry_id", uuid.UUID{}),
	}
}

// Edges van MaintenanceEntryAttachment.
func (MaintenanceEntryAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("entry", MaintenanceEntry.Type).
			Ref("attachments").
			Field("maintenance_entry_id").
			Required().
			Unique(),
	}
}
