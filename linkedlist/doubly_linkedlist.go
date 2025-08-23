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
	Head   *DoublyNode[T]
	Tail   *DoublyNode[T]
	Count  int
	equals func(a, b T) bool
}

// NewDoublyLinkedList creates and returns a new Doubly Linked List
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// GetElementAt returns an element that is at a specif position
func (dll *DoublyLinkedList[T]) GetElementAt(position int) (*DoublyNode[T], bool) {

	if position == 0 && position <= dll.Count {

		node := dll.Head
		for range position {
			node = node.Next
		}

		return node, true

	}

	var zero T
	return &DoublyNode[T]{Element: zero}, false

}

// Insert inserts an element at any position in the doubly linked list
func (dll *DoublyLinkedList[T]) Insert(element T, position int) bool {

	if position >= 0 && position <= dll.Count {

		node := NewDoublyNode(element)
		current := dll.Head

		if position == 0 {
			if dll.Head == nil {
				dll.Head = node
				dll.Tail = node
			} else {
				node.Next = dll.Head
				current.Prev = node
				dll.Head = node
			}
			dll.Count++
			return true

		}

		if position == dll.Count {
			current = dll.Tail
			current.Next = node
			node.Prev = current
			dll.Tail = node
			dll.Count++
			return true
		}

		previous, ok := dll.GetElementAt(position - 1)
		if !ok {
			return false
		}
		current = previous.Next
		node.Next = current
		previous.Next = node
		current.Prev = node
		node.Prev = previous

		dll.Count++
		return true
	}

	return false

}

// RemoveAt removes a specific element from the doubly list
func (dll *DoublyLinkedList[T]) RemoveAt(position int) (T, bool) {

	if position >= 0 && position < dll.Count {

		current := dll.Head

		if position == 0 {
			dll.Head = current.Next
			if dll.Count == 1 {
				dll.Tail = nil
			} else {
				dll.Head.Prev = nil
			}

			dll.Count--
			return current.Element, true

		}

		if position == dll.Count-1 {
			current = dll.Tail
			dll.Tail = current.Prev
			dll.Tail = nil
			dll.Count--
			return current.Element, true
		}

		current, ok := dll.GetElementAt(position)
		if !ok {
			var zero T
			return zero, false
		}

		previous := current.Prev
		previous.Next = current.Next
		current.Next.Prev = previous

		dll.Count--
		return current.Element, true

	}

	var zero T
	return zero, false

}
