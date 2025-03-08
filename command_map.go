package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type config struct {
	Next     string
	Previous string
}

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap(cnf *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cnf.Next != "" {
		url = cnf.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var locationResp locationAreaResponse
	err = json.Unmarshal(body, &locationResp)
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
