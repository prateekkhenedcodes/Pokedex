package main

import ("fmt")


func commandMap(cnf *config) error {
	NextUrl := cnf.Next
	locationResp, err := cnf.apiClient.ListLocation(&NextUrl)
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
