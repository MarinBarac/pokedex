package main

import "fmt"

func commandInspect(c *config, args ...string) error {
	name := args[0]
	pokemon, ok := c.pokemons[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}

	return nil
}
