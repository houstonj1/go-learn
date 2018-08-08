package main

import "fmt"

// Hello function
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("James"))
}
