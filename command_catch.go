package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cnf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name to catch it")
	}
	name := args[0]
	pokemon, err := cnf.pokeapiClient.ListPokemonAbility(name)
	if err != nil {
		return err
	}
	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if chance > 50 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("Now you can inspect this pokemon with inspect command")
		cnf.caughtPokemon[pokemon.Name] = pokemon
	}
	return nil
}
