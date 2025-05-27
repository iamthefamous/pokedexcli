package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func commandCatch(args []string, cfg *config) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	name := strings.ToLower(args[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	data, err := cfg.pokeapiClient.DetailPokemon(url)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(255)
	if chance > data.BaseExperience {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}

	fmt.Printf("%s was caught!\n", args[0])
	cfg.CaughtPokemon[name] = data
	return nil
}
