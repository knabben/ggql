package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/database"
	"github.com/spf13/cobra"
	"log"
)

var (
	sqlite, url string
	result      client.Schema
	ctx         = context.Background()
	variables   = map[string]interface{}{}
	uri         = fmt.Sprintf("file:%s?cache=shared&_fk=1", sqlite)
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

		// GraphQL client request and response serializer
		c := client.NewClient(url)
		c.GraphQL(client.BuildIntrospectionQuery(), variables, &result)

		// Start database to dump graph
		database := database.NewDatabase(uri)
		client, err := database.Connect()
		if err != nil {
			log.Fatalf("%v", err)
		}

		defer client.Close()
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		CreateObjectType(ctx, client, result)
	},
}

func CreateObjectType(ctx context.Context, client *ent.Client, result client.Schema) (o *ent.ObjectType, err error) {
	// start versioning schema
	// Find queryType

	//queryTypeName := result.Schema.QueryType.Name
	//
	//for _, types := range result.Schema.Types {
	//	if types.Name ==  queryTypeName {
	//
	//		for _, field := range types.Fields {
	//
	//			fieldType, err := client.FieldType.Create().SetName("field").Save(ctx)
	//			fmt.Println(field.Name)
	//			fmt.Println(field.Description)
	//		}
	//	}
	//}

	o, err = client.ObjectType.Create().SetName("name").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("object was created: ", o)
	return o, nil
}
