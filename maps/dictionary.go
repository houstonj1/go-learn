package maps

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotFound definition
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search method
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}