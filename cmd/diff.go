package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/pkg/client"

	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/ent/objecttype"

	"github.com/spf13/cobra"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var url string

func init() {
	diffCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8000/graphql/","GraphQL URI")
	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use: "diff",
	Short: "diff",
	Run: func(cmd *cobra.Command, args []string) {
		var result client.Schema
		var ctx = context.Background()
		variables := map[string]interface{}{}


		o, err := QueryObjectType(ctx, sqlClient)
		fmt.Println(o)
	},
}


func QueryObjectType(ctx context.Context, client *ent.Client) (*ent.ObjectType, error) {
	o := client.ObjectType.Query().Where(objecttype.NameEQ("name")).OnlyX(ctx)

	fields, err := o.QueryFields().All(ctx)
	for _, field := range fields {
		fmt.Println("FIELD: ", field)
	}

	if err != nil {
		return nil, fmt.Errorf("error querying")
	}
	log.Println("objecttype: ", o)

	return o, nil
}