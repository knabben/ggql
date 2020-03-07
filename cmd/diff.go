package cmd

import (
	"fmt"
	"github.com/knabben/ggql/pkg/client"
	"github.com/spf13/cobra"
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
	},
}