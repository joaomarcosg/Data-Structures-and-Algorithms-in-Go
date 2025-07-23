package deque

type Deque[T any] struct {
	items []T
}

// NewDeque creates and returns a new deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

// AddFront adds a new element to the deque
func (dq *Deque[T]) AddFront(element T) {
	dq.items = append([]T{element}, dq.items...)
}
