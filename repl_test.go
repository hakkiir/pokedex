package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello world, heres more words",
			expected: []string{"hello", "world,", "heres", "more", "words"},
		},
		{
			input:    "  Gello Forld, let's MAKE this a longer Sentence  ",
			expected: []string{"gello", "forld,", "let's", "make", "this", "a", "longer", "sentence"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len does not match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Println(word)
			fmt.Println((expectedWord))
			if word != expectedWord {
				t.Errorf("words dont match")
			}
		}
	}

}
