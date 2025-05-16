package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("You must provide a location name to explore")
	}

	if len(args) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	for _, explore := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", explore.Pokemon.Name)
	}

	return nil
}
