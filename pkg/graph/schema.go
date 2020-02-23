package graph

var s Schema

type Fields struct {
	Name string `json:"name"`
	Description string `json:"description"`
	IsDeprecated bool `json:"isDeprecated"`
}

type Schema struct {
	Schema struct {
		QueryType struct {
			Name   string   `json:"name"`
			Fields []Fields `json:"fields"`
		} `json:"queryType"`
	} `json:"__schema"`
}