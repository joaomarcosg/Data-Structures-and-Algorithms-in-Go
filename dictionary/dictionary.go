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

// KeyValue represents the key, value pair
type KeyValue[T comparable] struct {
	Key   string
	Value T
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

// KeyValues returns an array with all value pairs in the Dictionary
func (d *Dictionary[T]) KeyValues() []KeyValue[T] {
	valuePairs := make([]KeyValue[T], 0, len(d.Table))
	for key, value := range d.Table {
		valuePairs = append(valuePairs, KeyValue[T]{Key: key, Value: value})
	}
	return valuePairs
}

// Keys returns an array with all keys in the Dictionary
func (d *Dictionary[T]) Keys() []string {
	var keys []string
	for key := range d.Table {
		keys = append(keys, key)
	}
	return keys
}

// Values returns an array with all values in the Dictionary
func (d *Dictionary[T]) Values() []T {
	var values []T
	for _, value := range d.Table {
		values = append(values, value)
	}
	return values
}

// Size returns the numbers of elements in the Dictionary
func (d *Dictionary[T]) Size() int {
	return len(d.Table)
}
