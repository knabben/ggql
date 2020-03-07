package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Argument holds the schema definition for the Argument entity.
type Argument struct {
	ent.Schema
}

// Fields of the Argument.
func (Argument) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("type_kind"),
		field.String("type_name"),
	}
}

// Edges of the Argument.
func (Argument) Edges() []ent.Edge {
	return nil
}
