package cmd

import (
	"context"
	"fmt"

	c "github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/database"
	"github.com/knabben/ggql/pkg/graphql"
	"github.com/spf13/cobra"
	"log"
)

var (
	sqlite, url string
	result      graphql.Schema
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
		c.NewClient(url).GraphQL(graphql.BuildIntrospectionQuery(), variables, &result)

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
		_, err = database.CreateObjectType(ctx, result)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}