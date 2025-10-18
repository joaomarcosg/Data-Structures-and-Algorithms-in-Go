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

// loseloseHashCode is a simple hash function
func loseloseHashCode(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % 37
}

// HashCode returns the hash code for a given key
func (h *HashTable[T]) HashCode(key string) int {
	return loseloseHashCode(key)
}
