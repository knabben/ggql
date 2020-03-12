package graphql

import (
	"github.com/knabben/ggql/ent"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphqlCompareObjectType(t *testing.T) {
	var tests = []struct {
		source      ent.ObjectType
		destination ent.ObjectType
		result      bool
	}{
		{ent.ObjectType{}, ent.ObjectType{}, true},
		{ent.ObjectType{Name: "Query"}, ent.ObjectType{}, false},
		{ent.ObjectType{Name: "Query", Description: "description"}, ent.ObjectType{Name: "RootQuery", Description: ""}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, CompareObjectType(test.source, test.destination))
	}
}

func TestGraphqlCompareFields(t *testing.T) {

	var tests = []struct {
		source      []*ent.FieldType
		destination []*ent.FieldType
		added       []FieldError
		removed     []FieldError
	}{
		{
			[]*ent.FieldType{},
			[]*ent.FieldType{},
			[]FieldError{},
			[]FieldError{},
		},
		{
			[]*ent.FieldType{{Name: "name1"}, {Name: "name2"}},
			[]*ent.FieldType{{Name: "name2"}},
			[]FieldError{
				{Field: &ent.FieldType{Name: "name1"}, Message: ADDED_FIELD},
			},
			[]FieldError{},
		},
		{
			[]*ent.FieldType{{Name: "name1"}, {Name: "name2"}},
			[]*ent.FieldType{{Name: "name3"}},
			[]FieldError{
				{Field: &ent.FieldType{Name: "name1"}, Message: ADDED_FIELD},
				{Field: &ent.FieldType{Name: "name2"}, Message: ADDED_FIELD},
			},
			[]FieldError{
				{Field: &ent.FieldType{Name: "name3"}, Message: REMOVED_FIELD},
			},
		},
	}

	for _, test := range tests {
		added, removed := CompareFieldType(test.source, test.destination)
		assert.Equal(t, test.added, added)
		assert.Equal(t, test.removed, removed)
	}
}
