package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
    response := scanner.Text()
		cleaned := cleanInput(response)
		if len(cleaned) == 0 {
			continue
		}

		commandText := cleaned[0]
		ct, ok := registry[commandText]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := ct.callback()
		if err != nil {
			fmt.Println(err)
		}
		
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercaseText := strings.ToLower(trimmed)
	words := strings.Split(lowercaseText, " ")
	return words
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	registry := map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}
return registry
}