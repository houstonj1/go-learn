package di

import (
	"fmt"
	"io"
	"os"
)

// Greet function
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
