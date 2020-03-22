package graph

import (
	"fmt"
)

// Graph represents a graph
type Graph struct {
	nodes []Node
}

// Edge represents an edge in a directed graph
type Edge struct {
	src    int
	dest   int
	weight int
}

type edge struct {
	dest   *Node
	weight int
}

// Node represents a node in a graph
type Node struct {
	id    int
	neigh []edge
}

// GetSize returns the size of the graph
func (g *Graph) GetSize() int {
	return len(g.nodes)
}

// NewEmptyGraph creates an empty graph
func NewEmptyGraph() *Graph {
	return &Graph{}
}

// NewGraph creates a graph of a given size
func NewGraph(size int) *Graph {
	graph := &Graph{
		nodes: make([]Node, size),
	}

	for i := 0; i < size; i++ {
		graph.nodes[i].id = i
	}

	return graph
}

// AddEdge adds and edge (sourceID, destID)
func (g *Graph) AddEdge(sourceID, destID, weight int) error {
	if sourceID >= g.GetSize() {
		return fmt.Errorf("Invalid node %d", sourceID)
	}

	if destID >= g.GetSize() {
		return fmt.Errorf("Invalid node %d", destID)
	}

	g.nodes[sourceID].neigh = append(g.nodes[sourceID].neigh, edge{
		dest:   &g.nodes[destID],
		weight: weight,
	})
	return nil
}

// BuildGraph creates a directed graph from a list of edges
func BuildGraph(size int, edges []Edge) *Graph {
	graph := NewGraph(size)

	for _, edge := range edges {
		graph.AddEdge(edge.src, edge.dest, edge.weight)
	}

	return graph
}
