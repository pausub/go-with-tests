package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

// method with multiple return types
func (d Dictionary) Search(key string) (string, error) {
	// second return value indicates if value was found by key
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// d is a copy of a pointer to map, not copy of underlying structure
// so we must use "Dictionary" instead of "*Dictionary"
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
