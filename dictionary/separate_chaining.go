package dictionary

import (
	"joaomarcosg/Data-Structures-and-Algorithms-in-Go/linkedlist"
)

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

// loseloseHashCodeSeparateChaining is a simple hash function
func loseloseHashCodeSeparateChaining(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % 37
}

// HashCodeSeparateChaining returns the hash code for a given key
func (h *HashTableSeparateChaining[T]) HashCodeSeparateChaining(key string) int {
	return loseloseHashCodeSeparateChaining(key)
}

// Put adds a new item to the hash table
func (h *HashTableSeparateChaining[T]) Put(key string, value T) bool {
	if key != "" {
		position := h.HashCodeSeparateChaining(key)
		if h.Table[position] == nil {
			h.Table[position] = linkedlist.NewLinkedList(func(a, b KeyValuePair[T]) bool {
				return a.Key == b.Key
			})
		}
		h.Table[position].Push(KeyValuePair[T]{Key: key, Value: value})
		return true
	}

	return false
}

// Get returns a value found by key
func (h *HashTableSeparateChaining[T]) Get(key string) (T, bool) {
	var zero T
	if key == "" {
		return zero, false
	}

	position := h.HashCodeSeparateChaining(key)
	list := h.Table[position]
	if list != nil && !list.IsEmpty() {
		current := list.GetHead()
		for current != nil {
			if current.Element.Key == key {
				return current.Element.Value, true
			}
			current = current.Next
		}
	}
	return zero, false

}

// Remove removes an item from hash table
func (h *HashTableSeparateChaining[T]) Remove(key string) bool {
	if key != "" {
		position := h.HashCodeSeparateChaining(key)
		list := h.Table[position]

		if list != nil && !list.IsEmpty() {
			current := list.GetHead()
			for current != nil {
				if current.Element.Key == key {
					list.Remove(current.Element)
					if list.IsEmpty() {
						delete(h.Table, position)
					}
					return true
				}
				current = current.Next
			}
		}
		return false
	}

	return false
}
