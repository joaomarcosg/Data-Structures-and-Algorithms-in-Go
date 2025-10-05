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

// Set adds a new element in the Dictionary
func (d *Dictionary[T]) Set(key string, value T) bool {
	if key != "" {
		d.Table[key] = value
		return true
	}
	return false
}

// Remove removes an element from Dictionary using key as parameter
func (d *Dictionary[T]) Remove(key string) bool {
	if d.HasKey(key) {
		delete(d.Table, key)
		return true
	}
	return false
}

// Get returns a specific value from Dictionary using key as search parameter
func (d *Dictionary[T]) Get(key string) (T, bool) {
	value, ok := d.Table[key]
	return value, ok
}
