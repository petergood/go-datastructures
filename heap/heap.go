package heap

// Structure abstracts away the underlying data structure
type Structure interface {
	// Swap swaps elements i and j
	Swap(i, j int)
	// Push adds obj to the underlying data structure
	Push(obj interface{})
	// Pop removes the last element of the underying data structure
	Pop()
	// Compare returns true if the element at index i is <= than j, returns false otherwise
	Compare(i, j int) bool
	// Length returns the length of the underlying data structure
	Length() int
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func parent(index int) int {
	return index / 2
}

// Insert adds obj to the heap
func Insert(structure Structure, obj interface{}) {
	structure.Push(obj)
	curr := structure.Length() - 1

	for curr > 0 && structure.Compare(curr, parent(curr)) {
		structure.Swap(curr, parent(curr))
		curr = parent(curr)
	}
}

// RemoveMax removes the largest (according to relation specified in Compare) element from the heap
func RemoveMax(structure Structure) {
	structure.Swap(0, structure.Length()-1)
	structure.Pop()
	Heapify(structure, 0)
}

// Heapify performs the Heapify heap operation
func Heapify(structure Structure, index int) {
	minIndex := index
	l := left(index)
	r := right(index)

	if l < structure.Length() && structure.Compare(l, minIndex) {
		minIndex = l
	}

	if r < structure.Length() && structure.Compare(r, minIndex) {
		minIndex = r
	}

	if minIndex != index {
		structure.Swap(index, minIndex)
		Heapify(structure, l)
		Heapify(structure, r)
	}
}
