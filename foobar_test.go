package foobar

import "testing"

func TestFoo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "Hello, World!"},
		{"with name", "Alice", "Hello, Alice!"},
		{"with special chars", "Bob-123", "Hello, Bob-123!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Foo(tt.input)
			if result != tt.expected {
				t.Errorf("Foo(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBar(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "HELLO"},
		{"mixed case", "HeLLo", "HELLO"},
		{"empty string", "", ""},
		{"with spaces", "hello world", "HELLO WORLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Bar(tt.input)
			if result != tt.expected {
				t.Errorf("Bar(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFooBar(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "HELLO, WORLD!"},
		{"with name", "Alice", "HELLO, ALICE!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FooBar(tt.input)
			if result != tt.expected {
				t.Errorf("FooBar(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

