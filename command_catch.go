package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	pokemon, err := c.pokeapiClient.PokemonDetails(args[0])

	if err != nil {
		return err
	}
	randNum := rand.Intn(300)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if randNum > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)

		c.pokemons[pokemon.Name] = pokemon

		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon.Name)

	return nil
}
