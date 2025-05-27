package main

import "fmt"

func commandShowPokedex(args []string, cfg *config) error {
	if len(cfg.CaughtPokemon) < 1 {
		return fmt.Errorf("no pokemon caught yet!\n")
	}
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
