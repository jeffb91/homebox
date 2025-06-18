package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/schema/mixins"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("photo", "manual", "warranty", "attachment", "receipt").Default("attachment"),
		field.Bool("primary").Default(false),
		field.String("title").Default(""),
		field.String("path").Default(""),
<<<<<<< HEAD
		field.String("related_type").NotEmpty(),
		field.Int("related_id").Positive(),
=======
		// NIEUWE POLYMORPHIC VELDEN
		field.String("related_type").
			Comment("Type of related entity: items, maintenance_entries, incidents, reports"),
		field.UUID("related_id", uuid.UUID{}).
			Comment("ID of the related entity"),
>>>>>>> Attachments
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
<<<<<<< HEAD
	return nil
=======
	return []ent.Edge{
		// Verwijder de oude directe item edge
		// We gebruiken nu related_type + related_id voor polymorphic relations
	}
}

// Indexes of the Attachment.
func (Attachment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("related_type", "related_id"),
		index.Fields("related_type"),
	}
>>>>>>> Attachments
}
