package main

import (
	"errors"
	"fmt"
)

func commandExplore(args []string, cfg *config) error {
	if len(args) < 0 {
		return errors.New("not enough arguments")
	}
	areaName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	data, err := cfg.pokeapiClient.DetailLocation(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, pe := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}
	return nil
}
