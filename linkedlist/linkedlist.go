package linkedlist

// Node structure represents the linked list node
type Node[T any] struct {
	Element T
	Next    *Node[T]
}

func NewNode[T any](element T) *Node[T] {
	return &Node[T]{
		Element: element,
		Next:    nil,
	}
}

// Linkedlist structure with count representing the size of the list and its Head
type LinkedList[T any] struct {
	Head  *Node[T]
	count int
}

// NewLinkedList creates and returns a new Linked List
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Push adds a new element to the end of the linked list
func (ll *LinkedList[T]) Push(element T) {

	node := NewNode(element)

	if ll.Head == nil {
		ll.Head = node
		ll.count++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	ll.count++

}
