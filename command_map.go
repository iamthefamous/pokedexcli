package main

import (
	"errors"
	"fmt"
)

func commandMapf(args []string, cfg *config) error {
	location, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = location.Next
	cfg.prevLocationURL = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(args []string, cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	location, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = location.Next
	cfg.prevLocationURL = location.Previous
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
