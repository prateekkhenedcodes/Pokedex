package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemons{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemons{}, nil
		}

		return pokemonResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	pokemonResp := RespShallowPokemons{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil

}
