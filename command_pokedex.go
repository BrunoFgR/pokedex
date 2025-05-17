package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokemons := cfg.caughtPokemon
	if len(pokemons) == 0 {
		return errors.New("No Pokemon caught yet.")
	}

	fmt.Println("Your pokedex:")
	for pokemon, _ := range pokemons {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
