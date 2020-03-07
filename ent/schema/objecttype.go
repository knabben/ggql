package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// ObjectType holds the schema definition for the ObjectType entity.
type ObjectType struct {
	ent.Schema
}

// Fields of the ObjectType.
func (ObjectType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
	}
}

// Edges of the ObjectType.
func (ObjectType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fields", FieldType.Type),
	}
}
