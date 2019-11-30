package maps

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotFound definition
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search method
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add method
func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}
