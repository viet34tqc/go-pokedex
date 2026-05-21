package main

import (
	"go-http-client/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 30 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
