package dictionary

// HashTable represents a hash table
type HashTable[T comparable] struct {
	Table map[int]ValuePair[T]
}

// NewHashTable creates and returns a new hash table
func NewHashTable[T comparable]() *HashTable[T] {
	return &HashTable[T]{
		Table: make(map[int]ValuePair[T]),
	}
}

// Value represents the key, value pair
type ValuePair[T comparable] struct {
	Key   string
	Value T
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
