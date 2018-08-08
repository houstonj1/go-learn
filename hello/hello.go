package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("James", ""))
}
