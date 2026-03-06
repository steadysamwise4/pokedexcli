package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercaseText := strings.ToLower(trimmed)
	words := strings.Split(lowercaseText, " ")
	fmt.Println("Original:", text)
	fmt.Println("Lowercase:", lowercaseText)
	fmt.Println(words)
	return words
}