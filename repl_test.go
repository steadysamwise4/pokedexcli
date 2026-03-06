package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input string
	expected []string
	}{
			{
				input: " hello world  ",
				expected: []string{"hello", "world"},
			},
			{
				input: "Charmander Bulbasaur PIKACHU",
				expected: []string{"charmander", "bulbasaur", "pikachu"},
			},
			{
				input: "This is a POKEDEX",
				expected: []string{"this", "is", "a", "pokedex"},
			},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test Failed - expected slice length: %d, actual %d", len(actual), len(c.expected))
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test Failed - expected word: %s, actual: %s", word, expectedWord)
			}
		}
	}
}
