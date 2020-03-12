package cmd

import (
	"context"
	"fmt"
	"github.com/knabben/ggql/pkg/database"
	"github.com/knabben/ggql/pkg/parser"
	"go.uber.org/zap"

	"github.com/knabben/ggql/pkg/graphql"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var (
	result                            graphql.Schema
	sourceSchemaFile, sourceSchemaUrl string
	destSchemaFile, destSchemaUrl     string
	ctx                               = context.Background()
)

func init() {
	diffCmd.Flags().StringVarP(
		&sourceSchemaFile, "source-file", "s", "", "Source file schema",
	)
	diffCmd.Flags().StringVarP(
		&sourceSchemaUrl, "source-url", "u", "", "Source URL schema",
	)

	diffCmd.Flags().StringVarP(
		&destSchemaFile, "dest-file", "d", "", "Destination file schema",
	)
	diffCmd.Flags().StringVarP(
		&destSchemaUrl, "dest-url", "e", "", "Destination URL schema",
	)

	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrape",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		uri := "file:%s?mode=memory&cache=shared&_fk=1"

		sourceDatabase, err := database.NewDatabase(fmt.Sprintf(uri, "source"))
		if err != nil {
			logger.Fatal(err.Error())
		}
		destDatabase, err := database.NewDatabase(fmt.Sprintf(uri, "destination"))
		if err != nil {
			logger.Fatal(err.Error())
		}

		sourceParser := parser.NewParser(sourceSchemaFile, sourceSchemaUrl)
		sourceResult, err := sourceParser.LoadResult()
		if err != nil {
			logger.Fatal(err.Error())
		}

		sourceObject, err := sourceDatabase.CreateObjectType(ctx, sourceResult)
		if err != nil {
			logger.Fatal(err.Error())
		}

		destParser := parser.NewParser(destSchemaFile, destSchemaUrl)
		destResult, err := destParser.LoadResult()
		if err != nil {
			logger.Fatal(err.Error())
		}

		destObject, err := destDatabase.CreateObjectType(ctx, destResult)
		if err != nil {
			logger.Fatal(err.Error())
		}

		sourceFields, _ := sourceObject.QueryFields().All(ctx)
		destFields, _ := destObject.QueryFields().All(ctx)

		added, removed := graphql.CompareFieldType(sourceFields, destFields)
		for _, n := range added {
			fmt.Println(n)
		}
		fmt.Println("----- removed")
		for _, n := range removed {
			fmt.Println(n)
		}
	},
}
