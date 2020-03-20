package heap

import (
	"reflect"
	"testing"
)

func validateHeap(heap *Heap) bool {
	for i := 0; i <= (heap.size+2)/2; i++ {
		l := heap.left(i)
		r := heap.right(i)
		if (l < heap.size && heap.arr[l] < heap.arr[i]) ||
			(r < heap.size && heap.arr[r] < heap.arr[i]) {
			return false
		}
	}

	return true
}

func TestHeapify(t *testing.T) {
	heap := &Heap{
		arr:        []int{5, 1, 2, 8, 3, 2, 6, 4, 1, 7},
		size:       10,
		comparator: func(a, b int) bool { return a <= b },
	}

	heap.Heapify(0)

	if heap.arr[0] > heap.arr[1] || heap.arr[0] > heap.arr[2] {
		t.Errorf("Invalid heapify result")
	}
}

func TestInsert(t *testing.T) {
	heap := NewHeap(10, MinHeap)
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(4)
	heap.Insert(1)

	if !reflect.DeepEqual(heap.arr, []int{1, 4, 10, 5, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("Incorrect heap, got %v", heap.arr)
	}
}

func TestExtractMax(t *testing.T) {
	heap := NewHeap(10, MinHeap)
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(100)

	max1, max2 := heap.ExtractMax(), heap.ExtractMax()

	if max1 != 1 || max2 != 4 {
		t.Errorf("Incorrect max heap value %d %d", max1, max2)
	}

	if heap.size != 4 {
		t.Errorf("Incorrect heap size %d", heap.size)
	}

	if !validateHeap(heap) {
		t.Errorf("Heap not valid %v", heap.arr)
	}
}

func TestMaxHeap(t *testing.T) {
	heap := NewHeap(10, MaxHeap)
	heap.Insert(100)
	heap.Insert(20)
	heap.Insert(10)
	heap.Insert(50)
	heap.Insert(1)
	heap.Insert(314)

	max := heap.GetMax()

	if max != 314 {
		t.Errorf("Incorrect max value %d", max)
	}
}
