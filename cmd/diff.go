package cmd

import (
	 "context"
	"fmt"
	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/ent/objecttype"
	"github.com/knabben/ggql/pkg/client"
	"github.com/spf13/cobra"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var url string

func init() {
	diffCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8000/graphql/","GraphQL URI")
	rootCmd.AddCommand(diffCmd)
}


// Example of spec introspection
//{
//	"data": {
//	"__schema": {
//		"queryType": {
//			"fields": [
//			{
//				"name": "allVessels",
//				"type": {
//					"kind": "OBJECT",
//					"name": "VesselsObjectType"
//				}
//			},
//


var diffCmd = &cobra.Command{
	Use: "diff",
	Short: "diff",
	Run: func(cmd *cobra.Command, args []string) {
		var result client.Schema
		variables := map[string]interface{}{}

		c := client.NewClient(url)
		c.GraphQL(client.BuildIntrospectionQuery(), variables, &result)

		// Find queryType
		for _, types := range result.Schema.Types {
			if (types.Name == result.Schema.QueryType.Name) {
				for _, field := range types.Fields {
					fmt.Println(field.Name)
					fmt.Println(field.Description)
				}
			}
		}

		sqlClient, err := ent.Open("sqlite3", "file:ent1?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatalf("Failed opening connection to sqlite: %v", err)
		}

		defer sqlClient.Close()
		if err := sqlClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		ctx := context.Background()
		CreateObjectType(ctx, sqlClient)
		o, err := QueryObjectType(ctx, sqlClient)
		fmt.Println(o)
	},
}

func CreateObjectType(ctx context.Context, client *ent.Client) (o *ent.ObjectType, err error) {
	o, err = client.ObjectType.Create().SetName("name").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("object was created: ", o)
	return o, nil
}

func QueryObjectType(ctx context.Context, client *ent.Client) (*ent.ObjectType, error) {
	o, err := client.ObjectType.Query().Where(objecttype.NameEQ("name")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying")
	}
	log.Println("objecttype: ", o)
	return o, nil
}