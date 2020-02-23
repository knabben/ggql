package graph

import (
	"fmt"
	"sort"
)

var s Schema

type Types struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Fields struct {
	Name string `json:"name"`
}

type Schema struct {
	Schema struct {
		QueryType struct {
			Name   string   `json:"name"`
			Fields []Fields `json:"fields"`
		} `json:"queryType"`
	} `json:"__schema"`
}

func BuildGraph(s Schema) {
	fmt.Println(s.Schema)
	fields := []string{}

	for _, t := range s.Schema.QueryType.Fields {
		fields = append(fields, t.Name)

	}

	sort.Strings(fields)
	numAncestors := 3

	//map[string]string{}
	for n := 0; n < len(fields); n += len(fields) / numAncestors {
		fmt.Println(fields[n : n+numAncestors][0])
	}
}
