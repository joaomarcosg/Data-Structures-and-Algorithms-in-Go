package dictionary

// Dictionary represents a dictionary
type Dictionary[T comparable] struct {
	Table map[string]T
}

// NewDictionary creates and returns a new dictionary
func NewDictionary[T comparable]() *Dictionary[T] {
	return &Dictionary[T]{
		Table: make(map[string]T),
	}
}
