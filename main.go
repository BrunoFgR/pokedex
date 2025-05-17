package main

import (
	"time"

	"github.com/BrunoFgR/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.New(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
