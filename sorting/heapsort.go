package sorting

import (
	"github.com/petergood/go-datastructures/heap"
)

const (
	// Asc sorts in ascending order
	Asc = heap.MinHeap
	// Desc sorts in descending order
	Desc = heap.MaxHeap
)

// Sort sorts array using heapsort
func Sort(arr []int, dir int) []int {
	h := heap.BuildHeap(arr, dir)
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		res[i] = h.ExtractMax()
	}

	return res
}
