package main

import (
	"fmt"
	"testing"
)

func testCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "       hello         world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			fmt.Print(len(actual), len(c.expected))
			t.Errorf("test failed")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("test failed")
			}
		}
		fmt.Print("OK")
	}

}

// func testExitCommand(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{
// 			input: "exit",
// 			expected: "Closing the Pokedex... Goodbye!",
// 		},
// 		{
// 			input: "EXIT",
// 			expected: "Closing the Pokedex... Goodbye!",
// 		},

// 		{
// 			input: "ExIt",
// 			expected: "Closing the Pokedex... Goodbye!",
// 		},
// 		{
// 			input: "eXiT",
// 			expected: "Closing the Pokedex... Goodbye!",
// 		},
// 		{
// 			input: "      exit      ",
// 			expected: "Closing the Pokedex... Goodbye!",
// 		},

// 	}
// }
