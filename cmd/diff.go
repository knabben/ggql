package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	c "github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/database"
	"github.com/knabben/ggql/pkg/graphql"
	"github.com/spf13/cobra"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqlite, file, url string
	result            graphql.Schema
	ctx               = context.Background()
	variables         = map[string]interface{}{}
)

func init() {
	scrapeCmd.Flags().StringVarP(&url, "url", "u", "", "GraphQL URI")
	scrapeCmd.Flags().StringVarP(&file, "file", "f", "", "Schema file")
	scrapeCmd.Flags().StringVarP(&sqlite, "sqlite", "s", "sqlite3", "SQLite3 database")

	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape",
	Run: func(cmd *cobra.Command, args []string) {
		var uri = fmt.Sprintf("file:%s?_fk=1", sqlite)

		loadResult()

		database := database.NewDatabase(uri)
		client, err := database.Connect()
		if err != nil {
			log.Fatalf("%v", err)
		}

		defer client.Close()

		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		_, err = database.CreateObjectType(ctx, result)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func loadResult() {
	if _, err := os.Stat(sqlite); os.IsExist(err) {
		log.Fatalf("A database already exists.")
	}

	// Hit the GraphQL endpoint.
	if url != "" {
		parseGraphQLRequest()

	}

	// Use file for schema parser
	if file != "" {
		parseJSONSchema()
	}

	if result.Schema.QueryType.Name == "" {
		log.Fatalf("Error parsing the schema.")
	}
}

// parseGraphQLRequest GraphQL endpoint hit and result dump.
func parseGraphQLRequest() {

	err := c.NewClient(url).GraphQL(graphql.BuildIntrospectionQuery(), variables, &result)
	if err != nil {
		log.Fatalf("Trying to fetch GraphQL endpoint %s", err)
	}

	if result.Schema.QueryType.Name == "" {
		log.Fatalf("Error trying to fetch schema.")
	}
}

// parseJSONSchema parsing the file and dumps the result.
func parseJSONSchema() error {
	schema, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("Trying to read file: %v", err)
	}

	gr := &c.GraphQLResponse{Data: &result}
	err = json.Unmarshal(schema, &gr)
	if err != nil {
		log.Fatalf("Trying to unmarshal result: %v", err)
	}

	return nil
}
