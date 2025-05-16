package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResponse.Next
	cfg.prevLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Printf("- %s\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Printf("- %s\n", loc.Name)
	}
	return nil
}
