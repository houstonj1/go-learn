package maps

// Dictionary type
type Dictionary map[string]string

// Search method
func (d Dictionary) Search(word string) string {
	return d[word]
}
