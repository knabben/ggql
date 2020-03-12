package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/pkg/database"
	"github.com/knabben/ggql/pkg/parser"
	"go.uber.org/zap"

	"github.com/knabben/ggql/pkg/graphql"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var (
	uri = "file:%s?mode=memory&cache=shared&_fk=1"

	source, destination string
	ctx                               = context.Background()
)

func init() {
	diffCmd.Flags().StringVarP(
		&source,
		"source",
		"s",
		"",
		"Source URI can be a JSON or URL",
	)
	diffCmd.Flags().StringVarP(
		&destination,
		"destination",
		"d",
		"",
		"Destination file schema",
	)

	rootCmd.AddCommand(diffCmd)
}

func StartInternalDatabases(logger *zap.Logger) (*database.Database, *database.Database) {
	sourceDatabase, err := database.NewDatabase(fmt.Sprintf(uri, "source"))
	if err != nil {
		logger.Fatal(err.Error())
	}
	destDatabase, err := database.NewDatabase(fmt.Sprintf(uri, "destination"))
	if err != nil {
		logger.Fatal(err.Error())
	}

	return sourceDatabase, destDatabase
}

func createObject(logger *zap.Logger, url string, db *database.Database) *ent.ObjectType {
	parser := parser.NewParser(url)
	result, err := parser.LoadResult()
	if err != nil {
		logger.Fatal(err.Error())
	}

	object, err := db.CreateObjectType(ctx, result)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return object
}
var diffCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync() // nolint: errcheck

		sourceDatabase, destDatabase := StartInternalDatabases(logger)

		sourceObject := createObject(logger, source, sourceDatabase)
		destObject := createObject(logger, destination, destDatabase)

		sourceFields, _ := sourceObject.QueryFields().All(ctx)
		destFields, _ := destObject.QueryFields().All(ctx)

		errors := graphql.CompareFieldType(sourceFields, destFields)
		graphql.ParsePrintErrors(errors)
	},
}
