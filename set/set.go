package set

// Set represents a set
type Set[T comparable] struct {
	Items map[T]bool
}

// New creates and returns a new set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		Items: make(map[T]bool),
	}
}

// Has returns true if the element is in the set, and false otherwise
func (st *Set[T]) Has(element T) bool {
	return st.Items[element]
}

// Add adds a new element in the set
func (st *Set[T]) Add(element T) bool {
	if !st.Has(element) {
		st.Items[element] = true
		return true
	}

	return false
}
