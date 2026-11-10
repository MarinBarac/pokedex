package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"trailing empty space": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		"empty input": {
			input:    "",
			expected: []string{},
		},
		"spaces only": {
			input:    " ",
			expected: []string{},
		},
		"mixed whitespace only": {
			input:    " \t\n\r\v\f ",
			expected: []string{},
		},
		"single word": {
			input:    "hello",
			expected: []string{"hello"},
		},
		"multiple words": {
			input:    "hello world again",
			expected: []string{"hello", "world", "again"},
		},
		"leading whitespace": {
			input:    "  hello world",
			expected: []string{"hello", "world"},
		},
		"trailing whitespace": {
			input:    "hello world  ",
			expected: []string{"hello", "world"},
		},
		"repeated spaces": {
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		"tabs between words": {
			input:    "hello\tworld",
			expected: []string{"hello", "world"},
		},
		"newlines between words": {
			input:    "hello\nworld",
			expected: []string{"hello", "world"},
		},
		"windows line endings": {
			input:    "hello\r\nworld",
			expected: []string{"hello", "world"},
		},
		"other whitespace separators": {
			input:    "one\rtwo\vthree\ffour",
			expected: []string{"one", "two", "three", "four"},
		},
		"uppercase input": {
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		"mixed capitalization": {
			input:    "HeLLo WoRLd",
			expected: []string{"hello", "world"},
		},
		"combined normalization": {
			input:    " \tHeLLo \n WORLD\r\n ",
			expected: []string{"hello", "world"},
		},
		"punctuation preserved": {
			input:    "Hello, WORLD!",
			expected: []string{"hello,", "world!"},
		},
		"digits and symbols preserved": {
			input:    "ABC123 foo-bar a_b",
			expected: []string{"abc123", "foo-bar", "a_b"},
		},
		"unicode lowercase": {
			input:    "ČĆŽŠĐ CAFÉ",
			expected: []string{"čćžšđ", "café"},
		},
		"unicode whitespace": {
			input:    "hello\u00a0world\u2003again",
			expected: []string{"hello", "world", "again"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				fmt.Printf("expected len: %#v, got: %#v\n", len(c.expected), len(actual))
				t.Fatalf("%s: expected: %#v, got: %#v", name, c.expected, actual)
			}

			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]

				errorMessage := fmt.Sprintf("%s: expected word on index %v to be: %#v, got: %#v", name, i, expectedWord, word)

				if word != expectedWord {
					t.Error(errorMessage)
					t.Fatal(errorMessage)
				}
			}
		})
	}
}
