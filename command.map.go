package main

import(
	"fmt"
	// "github.com/steadysamwise4/pokedexcli/internal/pokeapi"
)

func fetchMapLocations(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
			return fmt.Errorf("failed to fetch locations: %w", err)
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func fetchPrevMapLocations(cfg *config) error {
	if cfg.prevLocationsURL == nil {
    fmt.Println("you're on the first page")
    return nil
	}
	
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)

	if err != nil {
			return fmt.Errorf("failed to fetch locations: %w", err)
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}