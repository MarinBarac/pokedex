package main

import "fmt"

func commandExplore(conf *config, args ...string) error {
	pokemons, err := conf.pokeapiClient.ListPokemons(args[0])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range pokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
