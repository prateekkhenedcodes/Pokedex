package main

import (
	"fmt"
)

func commandMapb(cnf *config) error {
	if cnf.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	previousLocationAreas, err := cnf.apiClient.ListLocation(cnf.Previous)
	if err != nil {
		return err
	}
	for _, previousLocationArea := range previousLocationAreas.Results {
		fmt.Println(previousLocationArea.Name)
	}
	return nil
}
