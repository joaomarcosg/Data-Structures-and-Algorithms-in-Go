package stacks

type Stack[T any] struct {
	items []T
}

// Newstack creates and returns a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(element T) {
	s.items = append(s.items, element)
}

// Pop removes and returns an element from the top of the stack
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

// Peek returns the top element of the stack without remove it
func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the numbers of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
