package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// FieldType holds the schema definition for the FieldType entity.
type FieldType struct {
	ent.Schema
}

// Fields of the FieldType.
func (FieldType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the FieldType.
func (FieldType) Edges() []ent.Edge {
	return nil
}
