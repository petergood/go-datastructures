package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgesParsing(t *testing.T) {
	edges := []Edge{Edge{0, 1, 1}, Edge{1, 2, 1}, Edge{2, 0, 2}, Edge{1, 3, 1}}
	size := 4

	graph := BuildGraph(size, edges)

	assert.Equal(t, len(graph.nodes[0].neigh), 1)
	assert.Equal(t, len(graph.nodes[1].neigh), 2)
	assert.Equal(t, len(graph.nodes[2].neigh), 1)
	assert.Equal(t, len(graph.nodes[3].neigh), 0)
	assert.Equal(t, graph.nodes[2].neigh[0].weight, 2)
}
