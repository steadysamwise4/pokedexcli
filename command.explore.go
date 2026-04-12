package main

import(
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No location name provided")
	}
	location, err := cfg.pokeapiClient.GetLocation(args[0])

	if err != nil {
			return fmt.Errorf("failed to fetch location: %w", err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}