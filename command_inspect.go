package main

import (
	"fmt"
)

func commandInspect(cnf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon to inspect it")
	}
	name := args[0]
	pokeInfo, ok := cnf.caughtPokemon[name]
	if !ok {
		fmt.Print("you have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n", pokeInfo.Name)
	fmt.Printf("Height: %d\n", pokeInfo.Height)
	fmt.Printf("Weight: %d\n", pokeInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeInfo.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	poketype := pokeInfo.Types
	for _, ptype := range poketype {
		fmt.Printf(" --%s\n", ptype.Type.Name)
	}

	return nil
}
