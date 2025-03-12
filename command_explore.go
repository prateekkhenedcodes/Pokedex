package main

import (
	"fmt"
)

func commandExplore(cnf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	name := args[0]
	pokemonResp, err := cnf.pokeapiClient.ListPokemons(name)
	if err != nil {
		// fmt.Print("here is the problem")
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	for _, val := range pokemonResp.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}

	return nil

}
