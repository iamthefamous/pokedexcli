package main

import (
	"bufio"
	"fmt"
	"github.com/iamthefamous/pokedexcli.git/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	CaughtPokemon   map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callBack(words[1:], cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callBack    func([]string, *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callBack:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callBack:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callBack:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callBack:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemon encountering area",
			callBack:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Through a Pokeball to catch a Pokemon",
			callBack:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows detailed information about a Pokemon",
			callBack:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows your pokedex",
			callBack:    commandShowPokedex,
		},
	}
}
