// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ArgumentsColumns holds the columns for the "arguments" table.
	ArgumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "type_kind", Type: field.TypeString},
		{Name: "type_name", Type: field.TypeString},
		{Name: "field_type_arguments", Type: field.TypeInt, Nullable: true},
	}
	// ArgumentsTable holds the schema information for the "arguments" table.
	ArgumentsTable = &schema.Table{
		Name:       "arguments",
		Columns:    ArgumentsColumns,
		PrimaryKey: []*schema.Column{ArgumentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "arguments_field_types_arguments",
				Columns: []*schema.Column{ArgumentsColumns[5]},

				RefColumns: []*schema.Column{FieldTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FieldTypesColumns holds the columns for the "field_types" table.
	FieldTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "deprecated_reason", Type: field.TypeString},
		{Name: "type_kind", Type: field.TypeString},
		{Name: "type_name", Type: field.TypeString},
		{Name: "object_type_fields", Type: field.TypeInt, Nullable: true},
	}
	// FieldTypesTable holds the schema information for the "field_types" table.
	FieldTypesTable = &schema.Table{
		Name:       "field_types",
		Columns:    FieldTypesColumns,
		PrimaryKey: []*schema.Column{FieldTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "field_types_object_types_fields",
				Columns: []*schema.Column{FieldTypesColumns[6]},

				RefColumns: []*schema.Column{ObjectTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ObjectTypesColumns holds the columns for the "object_types" table.
	ObjectTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// ObjectTypesTable holds the schema information for the "object_types" table.
	ObjectTypesTable = &schema.Table{
		Name:        "object_types",
		Columns:     ObjectTypesColumns,
		PrimaryKey:  []*schema.Column{ObjectTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArgumentsTable,
		FieldTypesTable,
		ObjectTypesTable,
	}
)

func init() {
	ArgumentsTable.ForeignKeys[0].RefTable = FieldTypesTable
	FieldTypesTable.ForeignKeys[0].RefTable = ObjectTypesTable
}
