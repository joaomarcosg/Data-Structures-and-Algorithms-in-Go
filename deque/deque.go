package deque

type Deque[T any] struct {
	items []T
}

// NewDeque creates and returns a new deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

// AddFront adds a new element at the front of the deque
func (dq *Deque[T]) AddFront(element T) {
	dq.items = append([]T{element}, dq.items...)
}

// AddBack adds a new element at the rear of the deque
func (dq *Deque[T]) AddBack(element T) {
	dq.items = append(dq.items, element)
}

// RemoveFront removes and returns an element from the front of the deque
func (dq *Deque[T]) RemoveFront() (T, bool) {
	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	frontElement := dq.items[0]
	dq.items = dq.items[1:]
	return frontElement, true
}

// RemoveBack removes and returns an element from the rear of the deque
func (dq *Deque[T]) RemoveBack() (T, bool) {
	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	rearElement := dq.items[len(dq.items)-1]
	dq.items = dq.items[:len(dq.items)-1]
	return rearElement, true
}

// FrontDeque returns the element from the front of the deque
func (dq *Deque[T]) FrontDeque() (T, bool) {
	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	return dq.items[0], true
}

// RearDeque returns the element from the rear of the deque
func (dq *Deque[T]) RearDeque() (T, bool) {
	if dq.IsEmpty() {
		var zero T
		return zero, false
	}

	return dq.items[len(dq.items)-1], true
}

// IsEmpty returns true if the deque is empty
func (dq *Deque[T]) IsEmpty() bool {
	return len(dq.items) == 0
}
