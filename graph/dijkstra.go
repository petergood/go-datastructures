package graph

import (
	"fmt"
	"math"

	"github.com/petergood/go-datastructures/heap"
)

type graphDist struct {
	minDist int
	nodeID  int
}

type graphDistHeap []graphDist

func (g *graphDistHeap) Swap(i, j int) {
	buf := (*g)[i]
	(*g)[i] = (*g)[j]
	(*g)[j] = buf
}

func (g *graphDistHeap) Push(obj interface{}) {
	*g = append(*g, obj.(graphDist))
}

func (g *graphDistHeap) Pop() {
	*g = (*g)[0 : len(*g)-1]
}

func (g *graphDistHeap) Compare(i, j int) bool {
	return (*g)[i].minDist <= (*g)[j].minDist
}

func (g *graphDistHeap) Length() int {
	return len(*g)
}

// ShortestPathWeight computes the weight of the shortest
// path between source and dest nodes
func ShortestPathWeight(graph *Graph, source, dest int) (int, []int, error) {
	// TODO: switch to Fibonacci heap
	graphDistHeap := &graphDistHeap{}
	shortestDists := make([]int, graph.GetSize())
	prev := make([]int, graph.GetSize())
	vis := make(map[int]struct{})

	for _, node := range graph.nodes {
		minDist := math.MaxInt32
		if node.id == source {
			minDist = 0
		}
		shortestDists[node.id] = minDist
		prev[node.id] = -1
	}

	heap.Insert(graphDistHeap, graphDist{
		minDist: 0,
		nodeID:  source,
	})

	for len(vis) != graph.GetSize() && graphDistHeap.Length() > 0 {
		gd := (*graphDistHeap)[0]
		heap.RemoveMax(graphDistHeap)
		vis[gd.nodeID] = struct{}{}

		for _, neigh := range graph.nodes[gd.nodeID].neigh {
			if gd.minDist+neigh.weight < shortestDists[neigh.dest.id] {
				shortestDists[neigh.dest.id] = gd.minDist + neigh.weight
				prev[neigh.dest.id] = gd.nodeID
				_, pres := vis[neigh.dest.id]
				if !pres {
					heap.Insert(graphDistHeap, graphDist{
						minDist: shortestDists[neigh.dest.id],
						nodeID:  neigh.dest.id,
					})
				}
			}
		}
	}

	if shortestDists[dest] == math.MaxInt32 {
		return 0, nil, fmt.Errorf("No path exists between %d and %d", source, dest)
	}

	path := []int{}
	curr := dest
	for curr != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	for i := 0; i < len(path)/2; i++ {
		opp := len(path) - i - 1
		path[i], path[opp] = path[opp], path[i]
	}

	return shortestDists[dest], path, nil
}
