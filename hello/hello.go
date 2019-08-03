package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns "Hello, " + name
// Separated concerns to make it easier to test
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprint(englishHelloPrefix, name)
}

func main() {
	name := ""
	fmt.Println(Hello(name))
}
