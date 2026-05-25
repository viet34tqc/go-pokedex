package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	chanceToCatch := rand.Intn(pokemon.BaseExperience)
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	if chanceToCatch > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("You caught %s!\n", pokemon.Name)
	return nil
}
