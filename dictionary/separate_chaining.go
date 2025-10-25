package dictionary

import "joaomarcosg/Data-Structures-and-Algorithms-in-Go/linkedlist"

// HashTableSeparateChaining represents a hash table with separate chaining technique
type HashTableSeparateChaining[T comparable] struct {
	Table map[int]*linkedlist.LinkedList[KeyValuePair[T]]
}

// NewHashTable creates and returns a new hash table with separate chaining technique
func NewHashTableSeparateChaining[T comparable]() *HashTableSeparateChaining[T] {
	return &HashTableSeparateChaining[T]{
		Table: make(map[int]*linkedlist.LinkedList[KeyValuePair[T]]),
	}
}

// KeyValuePair represents the key, value pair
type KeyValuePair[T comparable] struct {
	Key   string
	Value T
}
