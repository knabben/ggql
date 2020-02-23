package main

import (
	client2 "github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/graph"
)

var document = `query {
  __schema {
    queryType {
      fields {
        name
        description
        isDeprecated
      }
    }
  }
}`

func main() {
	var s graph.Schema

	client := client2.NewClient()
	client.GraphQL(document, map[string]interface{}{}, &s)

	graph.BuildGraph(s)
}
