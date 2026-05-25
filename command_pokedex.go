package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Println(pokemonName)
	}
	return nil
}
