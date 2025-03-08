package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		cleanedInput := cleanInput(input)
		value, exist := commands[cleanedInput[0]]
		if exist {
			err := value.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command\n")
			continue
		}

	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var pokedexConfig = &config{}
var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func() error {
				return commandExit(pokedexConfig)
			},
		},
		"help": {
			name:        "help",
			description: "Guide for Pokedex",
			callback: func() error {
				return commandHelp(commands, pokedexConfig)
			},
		},
		"map": {
			name:        "map",
			description: "20 location area of Pokemon world!",
			callback: func() error {
				return commandMap(pokedexConfig)
			},
		},
	}
}
