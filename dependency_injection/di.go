package di

import (
	"bytes"
	"fmt"
)

// Greet function
func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
