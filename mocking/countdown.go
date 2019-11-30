package mocking

import (
	"fmt"
	"io"
	"os"
)

// Countdown function
func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
