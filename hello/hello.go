package main

import "fmt"

const name = "Chris"

const englishHelloPrefix = "Hello, "

const spanish = "Spanish"

const spanishHelloPrefix = "Hola, "

const french = "French"

const frenchHelloPrefix = "Bonjour, "

// Hello returns "Hello, " + name
// Separated concerns to make it easier to test
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Using a named return initialises a variable to its zero value
// inside the function - this is why we can just `return` rather
// than `return prefix`
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello(name, ""))
}
