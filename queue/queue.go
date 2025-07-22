package queue

type Queue[T any] struct {
	items []T
}

// NewQueue creates and returns a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds a new element to the queue
func (q *Queue[T]) Enqueue(element T) {
	q.items = append(q.items, element)
}

// Dequeue removes and returns the first element added to the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true

}

// FrontQueue returns the element at the front of the queue
func (q *Queue[T]) FrontQueue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.items[0], true

}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the numbers of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}
