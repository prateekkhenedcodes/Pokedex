package main

import "fmt"

func commandPokedex(cnf *config, arg ...string) error {
	fmt.Println("your pokedex:")
	for _, name := range cnf.caughtPokemon {
		fmt.Printf(" - %s\n", name.Name)
	}
	return nil
}
