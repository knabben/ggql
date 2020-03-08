package cmd

import (
	"context"
	"fmt"

	"github.com/knabben/ggql/ent"
	"github.com/knabben/ggql/ent/objecttype"

	"github.com/spf13/cobra"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	rootCmd.AddCommand(diffCmd)
}

var diffCmd = &cobra.Command{
	Use: "diff",
	Short: "diff",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Diff")
		

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