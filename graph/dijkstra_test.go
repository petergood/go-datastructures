package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	edges := []Edge{
		Edge{0, 2, 1},
		Edge{0, 4, 5},
		Edge{0, 5, 3},
		Edge{4, 1, 1},
		Edge{5, 1, 6},
		Edge{2, 3, 7},
		Edge{4, 2, 2},
		Edge{5, 4, 1},
	}
	g := BuildGraph(6, edges)

	w, path, _ := ShortestPathWeight(g, 0, 1)

	if w != 5 {
		t.Errorf("Incorrect shortest path weight: %d", w)
	}

	if !reflect.DeepEqual([]int{0, 5, 4, 1}, path) {
		t.Errorf("Incorrect shortest path %v", path)
	}
}

func TestDijkstraWhenNoPathExists(t *testing.T) {
	edges := []Edge{
		Edge{0, 2, 1},
		Edge{0, 4, 5},
		Edge{0, 5, 3},
		Edge{1, 4, 1},
		Edge{1, 5, 6},
		Edge{2, 3, 7},
		Edge{2, 4, 2},
		Edge{4, 5, 1},
	}
	g := BuildGraph(6, edges)

	_, _, err := ShortestPathWeight(g, 0, 1)

	if err == nil {
		t.Errorf("Did not get error")
	}
}
