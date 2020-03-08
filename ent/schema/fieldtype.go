package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// FieldType holds the schema definition for the FieldType entity.
type FieldType struct {
	ent.Schema
}

// Fields of the FieldType.
func (FieldType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("deprecated_reason"),
		field.String("type_kind"),
		field.String("type_name"),
	}
}

// Edges of the FieldType.
func (FieldType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("arguments", Argument.Type),
	}
}
