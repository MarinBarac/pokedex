package main

import "fmt"

func commandHelp(conf *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	commands := conf.cmds

	for _, command := range commands {
		fmt.Printf("%s: %v\n", command.name, command.description)
	}

	return nil
}
