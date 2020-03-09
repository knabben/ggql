package graphql

import (
	"github.com/knabben/ggql/ent"
	"reflect"
)

var (
	REMOVED_FIELD = "REMOVED_FIELD"
	ADDED_FIELD   = "ADDED_FIELD"
)

type FieldError struct {
	Field   ent.FieldType
	Message string
}

func compareObjectType(source, destination ent.ObjectType) bool {
	return reflect.DeepEqual(source, destination)
}

func hasElement(source ent.FieldType, fieldtype []ent.FieldType) bool {
	for _, field := range fieldtype {
		if reflect.DeepEqual(source, field) {
			return true
		}
	}
	return false
}

func compareFieldType(source, destination []ent.FieldType) ([]FieldError, []FieldError) {
	var (
		removed = []FieldError{}
		added  = []FieldError{}
	)

	for _, sourceField := range source {
		if !hasElement(sourceField, destination) {
			added = append(added, FieldError{Field: sourceField, Message: ADDED_FIELD})
		}
	}

	for _, destField := range destination {
		if !hasElement(destField, source) {
			removed = append(removed, FieldError{Field: destField, Message: REMOVED_FIELD})
		}
	}

	return added, removed
}