package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/ent"
	c "github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/database"
	"github.com/spf13/cobra"
	"log"
)

var (
	sqlite, url string
	result      c.Schema
	ctx         = context.Background()
	variables   = map[string]interface{}{}
)

func init() {
	scrapeCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8000/graphql/", "GraphQL URI")
	scrapeCmd.Flags().StringVarP(&sqlite, "sqlite", "s", "sqlite3", "SQLite3 database")

	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape",
	Run: func(cmd *cobra.Command, args []string) {
		var uri = fmt.Sprintf("file:%s?_fk=1", sqlite)

		// GraphQL client request and response serializer
		c.NewClient(url).GraphQL(c.BuildIntrospectionQuery(), variables, &result)

		if result.Schema.QueryType.Name == "" {
			log.Fatalf("Error trying to fetch schema.")
		}

		// Start database to dump graph
		database := database.NewDatabase(uri)
		client, err := database.Connect()
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer client.Close()

		// create sqlite schema
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		// Dump schema on database
		o, err := CreateObjectType(ctx, client, result)
		if err != nil {
			log.Fatalf(err.Error())
		}

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
	},
}

// CreateFields return the list of fields
func CreateFields(client *ent.Client, types c.TypeDefinition) []*ent.FieldType {
	fields := []*ent.FieldType{}
	for _, field := range types.Fields {
		typeKind := field.Type["kind"].(string)
		typeValue := fmt.Sprintf("%s", field.Type["name"])

		fieldType, err := client.FieldType.
			Create().
			SetName(field.Name).
			SetDescription(field.Description).
			SetTypeKind(typeKind).
			SetDeprecatedReason(field.DeprecationReason).
			SetTypeName(typeValue).
			AddArguments(CreateArguments(client, field)...).
			Save(ctx)

		if err != nil {
			log.Println("Not adding this field %s: %s", field.Name, err)
			continue
		}
		fields = append(fields, fieldType)
	}

	return fields
}

// CreateObjectType convert the schema to a graph and dump in the database
func CreateObjectType(ctx context.Context, client *ent.Client, result c.Schema) (o *ent.ObjectType, err error) {
	for _, types := range result.Schema.Types {
		if types.Name == result.Schema.QueryType.Name {
			o, err = client.ObjectType.
				Create().
				SetName(types.Name).
				SetKind(types.Kind).
				SetDescription(types.Description).
				AddFields(CreateFields(client, types)...).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("failed creating objecttype: %v", err)
			}
		}
	}

	log.Println("object was created: ", o)
	return o, nil
}

// CreateArguments return the list of arguments of a field
func CreateArguments(client *ent.Client, field c.FieldDefinition) []*ent.Argument {
	arguments := []*ent.Argument{}

	for _, arg := range field.Args {
		argKind := arg.Type["kind"].(string)
		argValue := fmt.Sprintf("%s", arg.Type["name"])

		argumentType, err := client.Argument.
			Create().
			SetName(arg.Name).
			SetDescription(arg.Description).
			SetTypeKind(argKind).
			SetTypeName(argValue).
			Save(ctx)

		if err != nil {
			log.Println("Not adding this argument %s: %s", arg.Name, err)
			continue
		}
		arguments = append(arguments, argumentType)
	}

	return arguments
}
