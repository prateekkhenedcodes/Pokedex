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
		if stat.Stat.Name == "hp" {
			hp := stat.BaseStat
			fmt.Printf(" --hp: %d\n", hp)
		}
		if stat.Stat.Name == "attack" {
			attack := stat.BaseStat
			fmt.Printf(" --attack: %d\n", attack)
		}
		if stat.Stat.Name == "defense" {
			defense := stat.BaseStat
			fmt.Printf(" --defense: %d\n", defense)
		}
		if stat.Stat.Name == "special-attack" {
			special_attack := stat.BaseStat
			fmt.Printf(" --special-attack: %d\n", special_attack)
		}
		if stat.Stat.Name == "special-defense" {
			special_defense := stat.BaseStat
			fmt.Printf(" --soecial-defense: %d\n", special_defense)
		}
		if stat.Stat.Name == "speed" {
			speed := stat.BaseStat
			fmt.Printf(" --speed: %d\n", speed)
		}
	}
	fmt.Println("Types:")
	poketype := pokeInfo.Types
	for _, ptype := range poketype {
		fmt.Printf(" --%s\n", ptype.Type.Name)
	}

	return nil
}
