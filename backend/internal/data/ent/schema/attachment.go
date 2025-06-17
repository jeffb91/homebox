package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("related_type").NotEmpty(),
		field.UUID("related_id", uuid.UUID{}),
	}
}

// Edges of the Attachment.
//func (Attachment) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.From("item", Item.Type).
//			Ref("attachments").
//			Required().
//			Unique(),
//	}
//}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return nil // Geen edges meer, alles via related_type/related_id
}
