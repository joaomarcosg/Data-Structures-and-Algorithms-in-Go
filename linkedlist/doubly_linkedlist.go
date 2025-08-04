package linkedlist

// DoublyNode structure represents the doubly linked list node
type DoublyNode[T any] struct {
	Element T
	Prev    *DoublyNode[T]
	Next    *DoublyNode[T]
}

// NewDoublyNode creates a new doubly node
func NewDoublyNode[T any](element T) *DoublyNode[T] {
	return &DoublyNode[T]{
		Element: element,
		Prev:    nil,
		Next:    nil,
	}
}

// DoublyLinkedList structure representing the doubly linked list and its Head, Tail and Count (the size)
type DoublyLinkedList[T any] struct {
	Head  *DoublyNode[T]
	Tail  *DoublyNode[T]
	Count int
}

// NewDoublyLinkedList creates and returns a new Doubly Linked List
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}
