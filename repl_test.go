package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pokedex is much fun!",
			expected: []string{"pokedex", "is", "much", "fun!"},
		},
		{
			input:    "Understand how go works",
			expected: []string{"understand", "how", "go", "works"},
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}

		for i := range actual {
			if actual[i] != tc.expected[i] {
				t.Errorf("Expected %v, got %v", tc.expected[i], actual[i])
			}
		}
	}
}
