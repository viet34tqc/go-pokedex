package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	_, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", cfg.caughtPokemon[pokemonName].Height)
	fmt.Printf("Weight: %d\n", cfg.caughtPokemon[pokemonName].Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range cfg.caughtPokemon[pokemonName].Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range cfg.caughtPokemon[pokemonName].Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
