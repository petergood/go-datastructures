package heap

const (
	// MinHeap min heap
	MinHeap = iota
	// MaxHeap max heap
	MaxHeap = iota
)

// Comparator returns:
// - true, if a <= b
// - false, if a > b
type Comparator func(a, b int) bool

// Heap represents a heap
type Heap struct {
	arr        []int
	size       int
	comparator Comparator
}

// NewHeap creates a new heap
func NewHeap(maxSize int, heapType int) *Heap {
	heap := &Heap{
		arr:  make([]int, maxSize),
		size: 0,
	}

	if heapType == MinHeap {
		heap.comparator = func(a, b int) bool { return a <= b }
	} else {
		heap.comparator = func(a, b int) bool { return a >= b }
	}

	return heap
}

func (h *Heap) parent(node int) int {
	return node / 2
}

func (h *Heap) left(node int) int {
	return 2*node + 1
}

func (h *Heap) right(node int) int {
	return 2*node + 2
}

func (h *Heap) swap(i, j int) {
	buf := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = buf
}

// GetMax returns the largest/smallest element of the heap
func (h *Heap) GetMax() int {
	return h.arr[0]
}

// ExtractMax removes the largest/smallest element from the heap
func (h *Heap) ExtractMax() int {
	if h.size == 1 {
		return h.arr[0]
	}

	elem := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.Heapify(0)
	return elem
}

// Insert inserts a value into the heap
func (h *Heap) Insert(value int) {
	h.arr[h.size] = value
	curr := h.size
	h.size++

	for curr > 0 && h.comparator(h.arr[curr], h.arr[h.parent(curr)]) {
		buf := h.arr[curr]
		h.arr[curr] = h.arr[h.parent(curr)]
		h.arr[h.parent(curr)] = buf
		curr = h.parent(curr)
	}
}

// Heapify performs transforms an array into a valid heap
func (h *Heap) Heapify(node int) {
	l := h.left(node)
	r := h.right(node)
	minIndex := node

	if l < h.size && h.comparator(h.arr[l], h.arr[minIndex]) {
		minIndex = l
	}

	if r < h.size && h.comparator(h.arr[r], h.arr[minIndex]) {
		minIndex = r
	}

	if minIndex != node {
		h.swap(minIndex, node)
		h.Heapify(l)
		h.Heapify(r)
	}
}

// BuildHeap creates a heap from a given array
func BuildHeap(a []int, heapType int) *Heap {
	heap := NewHeap(len(a), heapType)
	heap.arr = make([]int, len(a))
	copy(heap.arr, a)
	heap.size = len(a)

	for i := len(a) - 1; i >= 0; i-- {
		heap.Heapify(i)
	}

	return heap
}
