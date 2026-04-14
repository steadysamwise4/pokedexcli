package main

import(
	"fmt"
	// "github.com/steadysamwise4/pokedexcli/internal/pokeapi"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon name provided")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
			return fmt.Errorf("failed to fetch pokemon: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchInt := rand.Intn(pokemon.BaseExperience)

	if catchInt > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	cfg.pokedex[name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}