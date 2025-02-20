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
			input:    "testing 1 2 3",
			expected: []string{"testing", "1", "2", "3"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("incorrect number of words: actual %v expected %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Fatalf("actual: %s did not match expected: %s", word, expectedWord)
			}
		}
	}
}
