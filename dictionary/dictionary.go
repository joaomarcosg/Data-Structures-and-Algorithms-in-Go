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

// Haskey returns true if key is present in the dictionary, and false otherwise
func (d *Dictionary[T]) HasKey(key string) bool {
	_, ok := d.Table[key]
	return ok
}
