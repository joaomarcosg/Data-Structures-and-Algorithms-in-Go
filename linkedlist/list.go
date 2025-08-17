package linkedlist

type List[T any] interface {
	Push(element T)
	RemoveAt(position int) (T, bool)
	GetElementAt(position int) (T, bool)
	Insert(element T, position int) bool
	IndexOf(element T) int
	Remove(element T) bool
	Size() int
	IsEmpty() bool
	GetHead() *T
}
