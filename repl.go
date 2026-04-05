package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"github.com/steadysamwise4/pokedexcli/internal/pokeapi"
)

type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
}

func startRepl(cfg *config) {
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
		err := ct.callback(cfg)
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
	callback func(*config) error
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
	"map": {
		name: "map",
		description: "Displays next 20 map locations",
		callback: fetchMapLocations,
	},
	"mapb": {
		name: "mapb",
		description: "Displays previous 20 map locations",
		callback: fetchPrevMapLocations,
	},
}
return registry
}