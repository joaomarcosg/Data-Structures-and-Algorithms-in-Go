package stacks

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(element T) {
	s.items = append(s.items, element)
}
