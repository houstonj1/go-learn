package main

import "fmt"

const helloPrefix = "Hello, "

// Hello function
func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("James"))
}
