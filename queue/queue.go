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
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true

}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
