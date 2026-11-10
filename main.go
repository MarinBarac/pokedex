package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)

	appConfig := config{
		cmds:          getCommands(),
		pokeapiClient: pokeapiClient,
		pokemons:      map[string]pokeapi.RespShallowPokemon{},
	}

	startRepl(&appConfig)
}
