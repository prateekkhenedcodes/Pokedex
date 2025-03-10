package main

import (
	"fmt"
)

func commandMapb(cnf *config) error {
	previousUrl := cnf.Previous
	previousLocationAreas, err := cnf.apiClient.ListLocation(&previousUrl)
	if err != nil {
		return nil 
	}
	for _, previousLocationArea := range previousLocationAreas.Results {
		fmt.Println(previousLocationArea.Name)
	}
	return nil
}
