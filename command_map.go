package main

import ("fmt")


func commandMap(cnf *config) error {
	nextUrl := cnf.Next
	locationResp, err := cnf.apiClient.ListLocation(&nextUrl)
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
