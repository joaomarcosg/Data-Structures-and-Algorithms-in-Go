package dictionary

// HashTable represents a hash map
type HashTable[T comparable] struct {
	Table map[string]T
}

// NewHashTable creates and returns a new hash map
func NewHashTable[T comparable]() *HashTable[T] {
	return &HashTable[T]{
		Table: make(map[string]T),
	}
}
