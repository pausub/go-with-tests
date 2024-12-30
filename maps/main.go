package main

import "errors"

type Dictionary map[string]string

// grouping of package visible variables
var (
	ErrNotFound        = errors.New("word not found")
	ErrWordExists      = errors.New("word already exists")
	ErrWordDoesntExist = errors.New("word doenst exist")
)

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
func (d Dictionary) Add(key, value string) error {
	_, found := d[key]
	if !found {
		d[key] = value
		return nil
	}
	return ErrWordExists
}

func (d Dictionary) Update(key, value string) error {
	_, found := d[key]
	if found {
		d[key] = value
		return nil
	}
	return ErrWordDoesntExist
}

func (d Dictionary) Delete(key string) error {
	_, found := d[key]
	if found {
		delete(d, key)
		return nil
	}
	return ErrWordDoesntExist
}
