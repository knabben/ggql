package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	scrapeCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8000/graphql/","GraphQL URI")
	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use: "scrape",
	Short: "scrape",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient(url)
		c.GraphQL(client.BuildIntrospectionQuery(), variables, &result)


		defer sqlClient.Close()


		if err := sqlClient.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}


		CreateObjectType(ctx, sqlClient, result)

	},
}


func CreateObjectType(ctx context.Context, client *ent.Client, result client.Schema) (o *ent.ObjectType, err error) {
	// start versioning schema
	// Find queryType

	queryTypeName := result.Schema.QueryType.Name

	for _, types := range result.Schema.Types {
		if types.Name ==  queryTypeName {

			for _, field := range types.Fields {

				fieldType, err := client.FieldType.Create().SetName("field").Save(ctx)
				fmt.Println(field.Name)
				fmt.Println(field.Description)
			}
		}
	}

	o, err = client.ObjectType.Create().SetName("name").AddFields(fieldType).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("object was created: ", o)

	return o, nil
}