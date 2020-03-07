// Code generated by entc, DO NOT EDIT.

package fieldtype

const (
	// Label holds the string label denoting the fieldtype type in the database.
	Label = "field_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description vertex property in the database.
	FieldDescription = "description"
	// FieldIsDeprecated holds the string denoting the is_deprecated vertex property in the database.
	FieldIsDeprecated = "is_deprecated"
	// FieldTypeKind holds the string denoting the type_kind vertex property in the database.
	FieldTypeKind = "type_kind"
	// FieldTypeName holds the string denoting the type_name vertex property in the database.
	FieldTypeName = "type_name"

	// Table holds the table name of the fieldtype in the database.
	Table = "field_types"
	// ArgumentsTable is the table the holds the arguments relation/edge.
	ArgumentsTable = "arguments"
	// ArgumentsInverseTable is the table name for the Argument entity.
	// It exists in this package in order to avoid circular dependency with the "argument" package.
	ArgumentsInverseTable = "arguments"
	// ArgumentsColumn is the table column denoting the arguments relation/edge.
	ArgumentsColumn = "field_type_arguments"
)

// Columns holds all SQL columns for fieldtype fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIsDeprecated,
	FieldTypeKind,
	FieldTypeName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FieldType type.
var ForeignKeys = []string{
	"object_type_fields",
}
