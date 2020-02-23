package graph

import (
	"fmt"
	"sort"
)

type Graph struct {
	Nodes map[string]Fields
	Links map[string]map[string]Fields
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]Fields),
		Links: make(map[string]map[string]Fields),
	}
}

// Check if node already exists
func (g *Graph) NodeExists(name string) bool {
	if _, exists := g.Nodes[name]; exists {
		return true
	}
	return false
}

// AddNode insert a node in the graph
func (g *Graph) AddNode(objectType Fields) {
	name := objectType.Name
	if(!g.NodeExists(name)) {
		g.Nodes[name] = objectType
	}
}

// AddEdge
func (g *Graph) AddEdge(vertexName1, vertexName2 string) bool {
	if !g.NodeExists(vertexName1) || !g.NodeExists(vertexName2) {
		return false
	}
	vertex2 := g.Nodes[vertexName2]

	var exists bool
	if _, exists = g.Links[vertexName1]; !exists {
		g.Links[vertexName1] = make(map[string]Fields)
	}
	g.Links[vertexName1][vertexName2] = vertex2

	return true
}

func (g *Graph) DepthFirstSearch() {
	for root, m := range g.Links {
		for vertex := range m {
			fmt.Printf("Link %v -> %v \n", root, vertex)
		}
	}
}

func BuildGraph(s Schema) {
	g := NewGraph()

	fields := []Fields{}
	for i, t := range s.Schema.QueryType.Fields {
		fields = append(fields, t)
		fmt.Printf("Node: %d %s \n", i, t.Name)
		g.AddNode(t)
	}

	// Sort slices by alpha ordering.
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name })

	// Split the first level of ancestor of the graph.
	step := len(fields)/3

	// Add the other first level ancestors.
	for n := 0; n < len(fields); n += step {
		ancestors := fields[n:step + n]
		for _, field := range ancestors[1:] {
			firstName := ancestors[0].Name
			g.AddEdge(firstName, field.Name)
		}
	}

	// Traverse edges and do request.
	g.DepthFirstSearch()
}
