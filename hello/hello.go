package main

import "fmt"

// Hello returns "Hello, world"
// Separate concerns to make it easier to test
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	name := "Chris"
	fmt.Println(Hello(name))
}
