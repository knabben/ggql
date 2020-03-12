package graphql

import (
	"fmt"
	"github.com/knabben/ggql/ent"
	"reflect"
)

var (
	FIELD_ADDED   = "FIELD_ADDED"
	ARG_ADDED     = "ARG_ADDED"
	FIELD_REMOVED = "FIELD_REMOVED"
	ARG_REMOVED   = "ARG_REMOVED"
)

type FieldError struct {
	Field   interface{}
	Error   string
	Message string
}

func CompareObjectType(source, destination ent.ObjectType) bool {
	return reflect.DeepEqual(source, destination)
}

// CompareArgument used to compare two lists of a field argument
func CompareArguments(source, destination []*ent.Argument) []FieldError {
	sourceInterface := make([]interface{}, len(source))
	for i, d := range source {
		sourceInterface[i] = d
	}
	destInterface := make([]interface{}, len(destination))
	for i, d := range destination {
		destInterface[i] = d
	}

	return compareItems(sourceInterface, destInterface)
}

// CompareFieldType used to compare two lists of fields
func CompareFieldType(source, destination []*ent.FieldType) []FieldError {
	sourceInterface := make([]interface{}, len(source))
	for i, d := range source {
		sourceInterface[i] = d
	}
	destInterface := make([]interface{}, len(destination))
	for i, d := range destination {
		destInterface[i] = d
	}

	return compareItems(sourceInterface, destInterface)
}

// compareItems compare different fields and returns the response.
func compareItems(source, destination []interface{}) []FieldError {
	var (
		errors = []FieldError{}
	)

	for _, sourceField := range source {
		if !hasElement(sourceField, destination) {
			field := FieldError{Field: sourceField}
			switch sourceField := sourceField.(type) {
			case *ent.Argument:
				field.Error = ARG_ADDED

			case *ent.FieldType:
				field.Message = fmt.Sprintf("%s was added", sourceField.Name)
				field.Error = FIELD_ADDED
			}
			errors = append(errors, field)
		}
	}
	for _, destField := range destination {
		if !hasElement(destField, source) {
			field := FieldError{Field: destField}
			switch destField := destField.(type) {
			case *ent.Argument:
				field.Error = ARG_REMOVED
			case *ent.FieldType:
				field.Message = fmt.Sprintf("%s was removed", destField.Name)
				field.Error = FIELD_REMOVED
			}
			errors = append(errors, field)
		}
	}

	return errors
}

func hasElement(source interface{}, fields []interface{}) bool {
	for _, field := range fields {
		switch field.(type) {
		case *ent.FieldType:
			if field.(*ent.FieldType).Name == source.(*ent.FieldType).Name {
				return true
			}
		case *ent.Argument:
			if field.(*ent.Argument).Name == source.(*ent.Argument).Name {
				return true
			}
		}
	}
	return false
}

func ParsePrintErrors(errors []FieldError) {
	for _, n := range errors {
		fmt.Printf("%s - %s\n", n.Error, n.Message)
	}
}
