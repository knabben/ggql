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

func TestGraphQLCompareFields(t *testing.T) {
	var tests = []struct {
		source      []*ent.FieldType
		destination []*ent.FieldType
		errors      []FieldError
	}{
		{
			[]*ent.FieldType{},
			[]*ent.FieldType{},
			[]FieldError{},
		},
		{
			[]*ent.FieldType{{Name: "name1"}, {Name: "name2"}},
			[]*ent.FieldType{{Name: "name2"}},
			[]FieldError{
				{Field: &ent.FieldType{Name: "name1"}, Message: ADDED_FIELD},
			},
		},
		{
			[]*ent.FieldType{
				{Name: "name1"}, {Name: "name2"},
			},
			[]*ent.FieldType{{Name: "name3"}},
			[]FieldError{
				{Field: &ent.FieldType{Name: "name1"}, Message: ADDED_FIELD},
				{Field: &ent.FieldType{Name: "name2"}, Message: ADDED_FIELD},
				{Field: &ent.FieldType{Name: "name3"}, Message: REMOVED_FIELD},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.errors, CompareFieldType(test.source, test.destination))
	}
}

func TestGraphQLCompareArguments(t *testing.T) {
	var tests = []struct {
		source      []*ent.Argument
		destination []*ent.Argument
		errors      []FieldError
	}{
		{
			[]*ent.Argument{},
			[]*ent.Argument{},
			[]FieldError{},
		},
		{
			[]*ent.Argument{{Name: "name1"}, {Name: "name2"}},
			[]*ent.Argument{{Name: "name2"}},
			[]FieldError{
				{Field: &ent.Argument{Name: "name1"}, Message: ADDED_ARGUMENT},
			},
		},
		{
			[]*ent.Argument{
				{Name: "name1"}, {Name: "name2"},
			},
			[]*ent.Argument{{Name: "name3"}},
			[]FieldError{
				{Field: &ent.Argument{Name: "name1"}, Message: ADDED_ARGUMENT},
				{Field: &ent.Argument{Name: "name2"}, Message: ADDED_ARGUMENT},
				{Field: &ent.Argument{Name: "name3"}, Message: REMOVED_ARGUMENT},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.errors, CompareArguments(test.source, test.destination))
	}
}
