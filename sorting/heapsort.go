package sorting

import (
	"github.com/petergood/go-datastructures/heap"
)

const (
	// Asc sorts in ascending order
	Asc = heap.MinIntHeap
	// Desc sorts in descending order
	Desc = heap.MaxIntHeap
)

// Sort sorts array using heapsort
func Sort(arr []int, dir int) []int {
	h := heap.BuildIntHeap(arr, dir)
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		res[i] = h.ExtractMax()
	}

	return res
}
