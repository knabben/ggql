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
	Field   interface{}
	Message string
}

func compareObjectType(source, destination ent.ObjectType) bool {
	return reflect.DeepEqual(source, destination)
}

func hasElement(source interface{}, fieldtype []interface{}) bool {
	for _, field := range fieldtype {
		if reflect.DeepEqual(source, field) {
			return true
		}
	}
	return false
}

// compareFieldType used to compare two lists of fields
func compareFieldType(source, destination []ent.FieldType) ([]FieldError, []FieldError) {
	sourceInterface := make([]interface{}, len(source))
	for i, d := range source { sourceInterface[i] = d}

	destInterface := make([]interface{}, len(destination))
	for i, d := range destination { destInterface[i] = d }

	return compareItems(sourceInterface, destInterface)
}

//// compareArgument used to compare two lists of a field argument
//func compareArgument(source, destination []ent.Argument) ([]FieldError, []FieldError) {
//	sourceInterface := make([]interface{}, len(source))
//	for i, d := range source { sourceInterface[i] = d }
//
//	destInterface := make([]interface{}, len(destination))
//	for i, d := range destination { destInterface[i] = d }
//
//	return compareItems(sourceInterface, destInterface)
//}

// compareItems
func compareItems(source, destination []interface{}) ([]FieldError, []FieldError) {
	var (
		removed = []FieldError{}
		added = []FieldError{}
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
