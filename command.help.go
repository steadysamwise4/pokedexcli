package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	registry := getCommands()
	for _, info := range registry {
		fmt.Printf("%s: %s\n", info.name, info.description)
	}
	fmt.Println()
	return nil
}