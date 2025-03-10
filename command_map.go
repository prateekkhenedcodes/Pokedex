package main

import (
	"fmt"
)

func commandMap(cnf *config) error {
	locationResp, err := cnf.apiClient.ListLocation(cnf.Next)
	if err != nil {
		return err
	}

	cnf.Next = locationResp.Next
	cnf.Previous = locationResp.Previous

	for _, locationName := range locationResp.Results {
		fmt.Println(locationName.Name)
	}
	return nil
}
