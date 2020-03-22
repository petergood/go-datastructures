package heap

const (
	// MinIntHeap min heap
	MinIntHeap = iota
	// MaxIntHeap max heap
	MaxIntHeap = iota
)

// Comparator returns:
// - true, if a <= b
// - false, if a > b
type Comparator func(a, b int) bool

// IntHeap represents a heap
type IntHeap struct {
	arr        []int
	size       int
	comparator Comparator
}

// NewIntHeap creates a new heap
func NewIntHeap(maxSize int, heapType int) *IntHeap {
	heap := &IntHeap{
		arr:  make([]int, maxSize),
		size: 0,
	}

	if heapType == MinIntHeap {
		heap.comparator = func(a, b int) bool { return a <= b }
	} else {
		heap.comparator = func(a, b int) bool { return a >= b }
	}

	return heap
}

func (h *IntHeap) parent(node int) int {
	return node / 2
}

func (h *IntHeap) left(node int) int {
	return 2*node + 1
}

func (h *IntHeap) right(node int) int {
	return 2*node + 2
}

func (h *IntHeap) swap(i, j int) {
	buf := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = buf
}

// GetMax returns the largest/smallest element of the heap
func (h *IntHeap) GetMax() int {
	return h.arr[0]
}

// ExtractMax removes the largest/smallest element from the heap
func (h *IntHeap) ExtractMax() int {
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
func (h *IntHeap) Insert(value int) {
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
func (h *IntHeap) Heapify(node int) {
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

// BuildIntHeap creates a heap from a given array
func BuildIntHeap(a []int, heapType int) *IntHeap {
	heap := NewIntHeap(len(a), heapType)
	heap.arr = make([]int, len(a))
	copy(heap.arr, a)
	heap.size = len(a)

	for i := len(a) - 1; i >= 0; i-- {
		heap.Heapify(i)
	}

	return heap
}
