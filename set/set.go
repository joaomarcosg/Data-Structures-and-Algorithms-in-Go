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

// Delete removes an element from the set
func (st *Set[T]) Delete(element T) bool {
	if st.Has(element) {
		delete(st.Items, element)
		return true
	}

	return false
}

// Size returns how many elements are in the set
func (st *Set[T]) Size() int {
	return len(st.Items)
}

// Values returns an array with all the elements in the set
func (st *Set[T]) Values() []T {
	keys := make([]T, 0, len(st.Items))
	for key := range st.Items {
		keys = append(keys, key)
	}
	return keys
}

// Union returns a new set with elements from both sets
func (st *Set[T]) Union(otherSet Set[T]) Set[T] {
	unionSet := New[T]()
	for i := range st.Items {
		unionSet.Add(i)
	}
	for i := range otherSet.Items {
		unionSet.Add(i)
	}
	return *unionSet
}
