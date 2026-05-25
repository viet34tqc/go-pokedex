package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location name provided")
	}
	locationName := args[0]
	fmt.Println("Exploring pastoria-city-area...")
	locationsResp, err := cfg.pokeapiClient.ListLocationsByName(locationName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, enc := range locationsResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil
}
