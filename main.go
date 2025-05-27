package main

import (
	"github.com/iamthefamous/pokedexcli.git/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
