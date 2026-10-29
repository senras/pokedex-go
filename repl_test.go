package main

import (
	"strings"
	"testing"
)

func TestCleanInput (t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "Hello, World!",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("len actual: %d, len expected: %d", len(actual), len(c.expected))
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, actual)
		}
		for i:= range actual{
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.ToLower(word) != expectedWord {
				t.Errorf("Expected %v, but got %v", expectedWord, word)
			}
		}
	}
}


func TestCleanInputEmpty (t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("len actual: %d, len expected: %d", len(actual), len(c.expected))
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, actual)
		}
		for i:= range actual{
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.ToLower(word) != expectedWord {
				t.Errorf("Expected %v, but got %v", expectedWord, word)
			}
		}
	}
}

func TestCleanInputSpecialCharacters (t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "Hello, World! This is a test.",
			expected: []string{"hello", "world", "this", "is", "a", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("len actual: %d, len expected: %d", len(actual), len(c.expected))
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, actual)
		}
		for i:= range actual{
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.ToLower(word) != expectedWord {
				t.Errorf("Expected %v, but got %v", expectedWord, word)
			}
		}
	}
}