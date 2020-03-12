package graphql

import (
	"fmt"
	"github.com/knabben/ggql/ent"
	"reflect"
)

var (
	ADDED_FIELD      = "ADDED_FIELD"
	ADDED_ARGUMENT   = "ADDED_ARGUMENT"
	REMOVED_FIELD    = "REMOVED_FIELD"
	REMOVED_ARGUMENT = "REMOVED_ARGUMENT"
)

type FieldError struct {
	Field   interface{}
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
			switch sourceField.(type) {
			case *ent.Argument:
				field.Message = ADDED_ARGUMENT
			case *ent.FieldType:
				field.Message = ADDED_FIELD
			}
			errors = append(errors, field)
		}
	}
	for _, destField := range destination {
		if !hasElement(destField, source) {
			field := FieldError{Field: destField}
			switch destField.(type) {
			case *ent.Argument:
				field.Message = REMOVED_ARGUMENT
			case *ent.FieldType:
				field.Message = REMOVED_FIELD
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
		fmt.Println(n)
	}
}
