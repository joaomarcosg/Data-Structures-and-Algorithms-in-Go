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

// Put adds a new item to the hash table
func (h *HashTable[T]) Put(key string, value T) bool {
	if key != "" {
		position := h.HashCode(key)
		h.Table[position] = ValuePair[T]{Key: key, Value: value}
		return true
	}
	return false
}

// Get returns a value found by key
func (h *HashTable[T]) Get(key string) (T, bool) {
	var zero T
	if key == "" {
		return zero, false
	}
	position := h.HashCode(key)
	valuePair, exists := h.Table[position]
	if !exists {
		return zero, false
	}
	return valuePair.Value, true
}

// Remove removes an item from hash table
func (h *HashTable[T]) Remove(key string) bool {
	if key == "" {
		return false
	}
	hash := h.HashCode(key)
	_, exists := h.Table[hash]
	if !exists {
		return false
	}
	delete(h.Table, hash)
	return true
}
