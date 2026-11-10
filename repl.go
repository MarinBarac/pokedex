package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type cliCommmand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	cmds             map[string]cliCommmand
	pokeapiClient    pokeapi.Client
	pokemons         map[string]pokeapi.RespShallowPokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func cleanInput(text string) []string {
	newText := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return newText
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := conf.cmds

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			continue
		}

		words := cleanInput(scanner.Text())
		firstWord := words[0]
		args := make([]string, 0)

		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := commands[firstWord]

		if ok == false {
			fmt.Println("Unknown command")
			commandHelp(conf, "")
		} else {
			command.callback(conf, args...)
		}

	}
}

func getCommands() map[string]cliCommmand {
	return map[string]cliCommmand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays the names of pokemons encountered on some location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Use this command to try catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect pokemon details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all your captured pokemons",
			callback:    commandPokedex,
		},
	}
}
