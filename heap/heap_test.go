package heap

import (
	"reflect"
	"testing"
)

type FloatHeap []float64

func (h *FloatHeap) Swap(i, j int) {
	buf := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = buf
}

func (h *FloatHeap) Push(obj interface{}) {
	*h = append(*h, obj.(float64))
}

func (h *FloatHeap) Pop() {
	*h = (*h)[0 : len(*h)-1]
}

func (h *FloatHeap) Compare(i, j int) bool {
	return (*h)[i] <= (*h)[j]
}

func (h *FloatHeap) Length() int {
	return len(*h)
}

func TestHeapify(t *testing.T) {
	hs := FloatHeap{3, 2, 1}

	Heapify(&hs, 0)

	if !reflect.DeepEqual(FloatHeap{1, 2, 3}, hs) {
		t.Errorf("Incorrect heap %v", hs)
	}
}

func TestInsertion(t *testing.T) {
	hs := FloatHeap{}

	Insert(&hs, 10.0)
	Insert(&hs, 3.0)
	Insert(&hs, 15.0)
	Insert(&hs, 1.0)

	if !reflect.DeepEqual(FloatHeap{1, 3, 15, 10}, hs) {
		t.Errorf("Incorrect heap %v", hs)
	}
}

func TestRemoveMax(t *testing.T) {
	hs := FloatHeap{}

	Insert(&hs, 10.0)
	Insert(&hs, 3.0)
	Insert(&hs, 15.0)
	Insert(&hs, 1.0)

	if hs[0] != 1 {
		t.Errorf("Incorrect min %f", hs[0])
	}

	RemoveMax(&hs)

	if hs.Length() != 3 {
		t.Errorf("Incorrect length %d", hs.Length())
	}

	if !reflect.DeepEqual(FloatHeap{3, 10, 15}, hs) {
		t.Errorf("Incorrect heap %v", hs)
	}

	RemoveMax(&hs)
	RemoveMax(&hs)
	RemoveMax(&hs)

	if hs.Length() != 0 {
		t.Errorf("Heap is not empty: %d", hs.Length())
	}
}
