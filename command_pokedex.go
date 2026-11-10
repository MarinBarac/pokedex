package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.pokemons) == 0 {
		fmt.Println("You didn't catch any pokemons yet")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range c.pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
