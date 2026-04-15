package main

import(
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// pokedex, ok := cfg.pokedex
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("No pokemon have been caught!")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}