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
	Head   *Node[T]
	count  int
	equals func(a, b T) bool
}

// NewLinkedList creates and returns a new Linked List
func NewLinkedList[T any](equals func(a, b T) bool) *LinkedList[T] {
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

// RemoveAt removes a specific element from the list
func (ll *LinkedList[T]) RemoveAt(position int) (T, bool) {

	if position >= 0 && position < ll.count {

		current := ll.Head

		if position == 0 {
			ll.Head = current.Next
			return current.Element, true
		}

		var previous *Node[T]

		for range position {
			previous = current
			current = current.Next
		}

		previous.Next = current.Next
		ll.count--
		return current.Element, true

	}

	var zero T
	return zero, false

}

// GetElementAt returns an element that is at a specif position
func (ll *LinkedList[T]) GetElementAt(position int) (*Node[T], bool) {

	if position >= 0 && position <= ll.count {

		node := ll.Head

		for range position {
			node = node.Next
		}

		return node, true

	}

	var zero T
	return &Node[T]{Element: zero}, false

}

// Insert inserts an element at any position in the list
func (ll *LinkedList[T]) Insert(element T, position int) bool {

	if position >= 0 && position >= ll.count {

		node := NewNode(element)

		if position == 0 {
			current := ll.Head
			node.Next = current
			ll.Head = node
			return true
		}

		previous, ok := ll.GetElementAt(position - 1)
		if !ok {
			return false
		}

		current := previous.Next
		node.Next = current
		previous.Next = node

		ll.count++
		return true
	}

	return false

}

// IndexOf returns a position of a given element
func (ll *LinkedList[T]) IndexOf(element T, equals func(a, b T) bool) int {

	current := ll.Head
	for i := 0; i < ll.count && current != nil; i++ {

		if equals(element, current.Element) {
			return i
		}

		current = current.Next

	}

	return -1

}
