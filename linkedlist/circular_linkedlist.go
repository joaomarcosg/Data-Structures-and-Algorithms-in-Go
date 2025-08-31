package linkedlist

// CircularNode structure represents the circular linked list node
type CircularNode[T any] struct {
	Element T
	Prev    *CircularNode[T]
	Next    *CircularNode[T]
}

// NewCircularNode creates a new circular node
func NewCircularNode[T any](element T) *CircularNode[T] {
	node := &CircularNode[T]{Element: element}
	node.Prev = node
	node.Next = node
	return node
}

// CircularLinkedList structure representing the circular linked list and its Head, Tail and Count (the size)
type CircularLinkedList[T any] struct {
	Head   *CircularNode[T]
	Tail   *CircularNode[T]
	Count  int
	equals func(a, b T) bool
}

// NewCircularLinkedList creates and returns a new Circular Linked List
func NewCircularLinkedList[T any](equals func(a, b T) bool) *CircularLinkedList[T] {
	return &CircularLinkedList[T]{}
}
