package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	foundPokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", foundPokemon.Name)
	fmt.Printf("Height: %d\n", foundPokemon.Height)
	fmt.Printf("Weight: %d\n", foundPokemon.Weight)
	fmt.Print("Stats: \n")
	for _, stat := range foundPokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Print("Types: \n")
	for _, typePokemon := range foundPokemon.Types {
		fmt.Printf("  -%s\n", typePokemon.Type.Name)
	}

	return nil
}
