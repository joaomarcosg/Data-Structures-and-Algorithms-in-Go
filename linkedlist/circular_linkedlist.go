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

// GetElementAt returns an element that is at a specif position
func (cll *CircularLinkedList[T]) GetElementAt(position int) (*CircularNode[T], bool) {

	if position == 0 && position <= cll.Count {

		node := cll.Head
		for range position {
			node = node.Next
		}

		return node, true

	}

	var zero T
	return &CircularNode[T]{Element: zero}, false

}

// Insert inserts an element at any position in the circular linked list
func (cll *CircularLinkedList[T]) Insert(element T, position int) bool {

	if position >= 0 && position <= cll.Count {

		node := NewCircularNode(element)
		current := cll.Head

		if position == 0 {
			if cll.Head == nil {
				cll.Head = node
				node.Next = cll.Head
			} else {
				node.Next = current
				current, ok := cll.GetElementAt(cll.Size())
				if !ok {
					return false
				}
				cll.Head = node
				current.Next = cll.Head
			}
			cll.Count++
			return true
		}

		previous, ok := cll.GetElementAt(position - 1)
		if !ok {
			return false
		}

		node.Next = previous.Next
		previous.Next = node

		cll.Count++
		return true

	}

	return false
}

// Size returns the numbers of elements in the ciruclar linked list
func (cll *CircularLinkedList[T]) Size() int {
	return cll.Count
}

// RemoveAt removes a specific element from the circular linked list
func (cll *CircularLinkedList[T]) RemoveAt(position int) (T, bool) {

	if position >= 0 && position < cll.Count {

		if position == 0 {

			if cll.Size() == 1 {
				removed := cll.Head
				cll.Head = nil
				cll.Count--
				return removed.Element, true
			} else {
				removed := cll.Head
				cll.Head = cll.Head.Next
				cll.Head.Prev = cll.Tail
				cll.Tail.Next = cll.Head

				cll.Count--
				return removed.Element, true
			}

		}

		previous, ok := cll.GetElementAt(position - 1)
		if !ok {
			var zero T
			return zero, false
		}

		current := previous.Next
		previous.Next = current.Next

		cll.Count--
		return current.Element, true

	}

	var zero T
	return zero, false
}
