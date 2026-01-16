package foobar

import (
	"fmt"
	"strings"
)

// Foo returns a greeting message with the given name.
// This is a sample function demonstrating basic string manipulation.
func Foo(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Bar transforms the input string to uppercase.
// This is a sample function demonstrating string transformation.
func Bar(input string) string {
	return strings.ToUpper(input)
}

// FooBar combines Foo and Bar operations.
// This demonstrates how functions can be composed.
func FooBar(name string) string {
	return Bar(Foo(name))
}
