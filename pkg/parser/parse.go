package parser

import (
	"encoding/json"
	c "github.com/knabben/ggql/pkg/client"
	"github.com/knabben/ggql/pkg/graphql"
	"io/ioutil"
	"strings"
)

type Parse struct {
	schemaFile, schemaUrl string
}

func NewParser(url string) *Parse {
	var parse = &Parse{schemaFile: "", schemaUrl: ""}

	if strings.HasPrefix(url, "http") {
		parse.schemaUrl = url
	} else {
		parse.schemaFile = url
	}

	return parse
}

// LoadResult choose from where to load the schema
func (p *Parse) LoadResult() (graphql.Schema, error) {
	// Hit the GraphQL endpoint.
	if p.schemaUrl != "" {
		return p.parseURLSchema()
	}

	// Use file for schema parser
	if p.schemaFile != "" {
		return p.parseFileSchema()
	}

	return graphql.Schema{}, nil
}

// parseGraphQLRequest Hits an endpoint and dumps the result
func (p *Parse) parseURLSchema() (graphql.Schema, error) {
	var (
		variables = map[string]interface{}{}
		result    = graphql.Schema{}
	)

	err := c.NewClient(p.schemaUrl).GraphQL(graphql.BuildIntrospectionQuery(), variables, &result)
	if err != nil {
		return result, err
	}

	if result.Schema.QueryType.Name == "" {
		return result, nil
	}

	return result, nil
}

// parseFileSchema parses a schema file and dumps the result.
func (p *Parse) parseFileSchema() (graphql.Schema, error) {
	var (
		result = graphql.Schema{}
	)

	schema, err := ioutil.ReadFile(p.schemaFile)
	if err != nil {
		return result, err
	}

	gr := &c.GraphQLResponse{Data: &result}

	err = json.Unmarshal(schema, &gr)
	if err != nil {
		return result, err
	}

	return result, nil
}