package dictionary

// HashTableLinearProbing represents a hash table with linear probing technique
type HashTableLinearProbing[T comparable] struct {
	Table map[int]ValuePair[T]
}

// NewHashTableLinearProbing creates and returns a new hash table with linear probing technique
func NewHashTableLinearProbing[T comparable]() *HashTable[T] {
	return &HashTable[T]{
		Table: make(map[int]ValuePair[T]),
	}
}

// PairKeyValue represents the key, value pair
type PairKeyValue[T comparable] struct {
	Key   string
	Value T
}

// loseloseHashLinearProbing is a simple hash function
func loseloseHashCodeLinearProbing(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % 37
}

// HashCodeLinearProbing returns the hash code for a given key
func (h *HashTableSeparateChaining[T]) HashCodeLinearProbing(key string) int {
	return loseloseHashCodeLinearProbing(key)
}
