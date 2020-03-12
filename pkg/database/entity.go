package database

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/pkg/graphql"
	"log"
)

var ctx = context.Background()

// DumpSchema traverse the graph
func DumpSchema(o ent.ObjectType) {
	fields, err := o.QueryFields().All(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, field := range fields {
		fmt.Println(field)

		args, _ := field.QueryArguments().All(ctx)
		for _, arg := range args {
			fmt.Println(arg)
		}
		fmt.Println("---")
	}
}

// CreateFields return the list of fields
func (d *Database) CreateFields(types graphql.TypeDefinition) []*ent.FieldType {
	fields := []*ent.FieldType{}

	for _, field := range types.Fields {
		typeKind := field.Type["kind"].(string)
		typeValue := fmt.Sprintf("%s", field.Type["name"])

		fieldType, err := d.client.FieldType.
			Create().
			SetName(field.Name).
			SetDescription(field.Description).
			SetTypeKind(typeKind).
			SetDeprecatedReason(field.DeprecationReason).
			SetTypeName(typeValue).
			AddArguments(d.CreateArguments(field)...).
			Save(ctx)

		if err != nil {
			fmt.Println(fmt.Sprintf("Not adding this field %s: %s", field.Name, err))
			continue
		}
		fields = append(fields, fieldType)
	}

	return fields
}

// CreateObjectType convert the schema to a graph and dump in the database
func (d *Database) CreateObjectType(ctx context.Context, result graphql.Schema) (o *ent.ObjectType, err error) {
	for _, types := range result.Schema.Types {
		if types.Name == result.Schema.QueryType.Name {
			o, err = d.client.ObjectType.
				Create().
				SetName(types.Name).
				SetKind(types.Kind).
				SetDescription(types.Description).
				AddFields(d.CreateFields(types)...).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("failed creating objecttype: %v", err)
			}
		}
	}
	return o, nil
}

// CreateArguments return the list of arguments of a field
func (d *Database) CreateArguments(field graphql.FieldDefinition) []*ent.Argument {
	arguments := []*ent.Argument{}

	for _, arg := range field.Args {
		argKind := arg.Type["kind"].(string)
		argValue := fmt.Sprintf("%s", arg.Type["name"])
		argumentType, err := d.client.Argument.
			Create().
			SetName(arg.Name).
			SetDescription(arg.Description).
			SetTypeKind(argKind).
			SetTypeName(argValue).
			Save(ctx)

		if err != nil {
			fmt.Println(fmt.Sprintf("Not adding this argument %s: %s", arg.Name, err))
			continue
		}
		arguments = append(arguments, argumentType)
	}

	return arguments
}

//func (d *Database) QueryObjectTypeFields(ctx context.Context, objectName string) ([]*ent.FieldType, error) {
//	fields, err := d.client.ObjectType.Query().Where(objecttype.NameEQ(objectName)).QueryFields().All(ctx)
//	if err != nil {
//		return nil, err
//	}
//	log.Println("objecttype fields:", fields)
//	return fields, nil
//}
