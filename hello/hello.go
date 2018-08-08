package main

import "fmt"

const helloPrefix = "Hello, "

// Hello function
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("James"))
}
