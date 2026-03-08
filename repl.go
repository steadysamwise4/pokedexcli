package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
    response := scanner.Text()
		cleaned := cleanInput(response)
		if len(cleaned) == 0 {
			continue
		}

		firstWord := cleaned[0]

		fmt.Printf("Your command was: %s\n", firstWord)
		
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercaseText := strings.ToLower(trimmed)
	words := strings.Split(lowercaseText, " ")
	return words
}