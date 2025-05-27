package main

import "fmt"

func commandInspect(args []string, cfg *config) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}
	p, ok := cfg.CaughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
