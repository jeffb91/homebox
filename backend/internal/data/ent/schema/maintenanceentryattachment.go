package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MaintenanceEntryAttachment holds the schema definition for the attachment.
type MaintenanceEntryAttachment struct {
	ent.Schema
}

func (MaintenanceEntryAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").NotEmpty(),
		field.String("filepath").NotEmpty(),
		field.Time("uploaded_at").Default(time.Now),
	}
}

func (MaintenanceEntryAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("entry", MaintenanceEntry.Type).
			Ref("attachments").
			Required().
			Unique(),
	}
}
