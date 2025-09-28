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
func (st *Set[T]) Union(otherSet *Set[T]) *Set[T] {
	unionSet := New[T]()
	for i := range st.Items {
		unionSet.Add(i)
	}
	for i := range otherSet.Items {
		unionSet.Add(i)
	}
	return unionSet
}

// Intersection returns a new set with elements present in both sets
func (st *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	intersectionSet := New[T]()
	values := st.Items
	otherValues := otherSet.Items
	biggerSet := values
	smallerSet := otherValues
	if len(otherValues)-len(values) > 0 {
		biggerSet = otherValues
		smallerSet = values
	}
	for item := range smallerSet {
		if biggerSet[item] {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

// Difference returns a new set with elements present in the first set but not in the second
func (st *Set[T]) Difference(otherSet *Set[T]) *Set[T] {
	differenceSet := New[T]()
	values := st.Items
	for item := range values {
		if !otherSet.Has(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

// IsSubsetOf returns true if a set is a subset of another
func (st *Set[T]) IsSubsetOf(otherSet *Set[T]) bool {
	if len(st.Items) > len(otherSet.Items) {
		return false
	}
	isSubset := true
	values := st.Items
	for item := range values {
		if !otherSet.Has(item) {
			isSubset = false
			return false
		}
		return true
	}
	return isSubset
}
