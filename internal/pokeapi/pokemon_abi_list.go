package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonAbility(pokemon_name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon_name

	if val, ok := c.cache.Get(url); ok {
		RespPokemonAbility := Pokemon{}
		err := json.Unmarshal(val, &RespPokemonAbility)
		if err != nil {
			return Pokemon{}, err
		}

		return RespPokemonAbility, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print("req error")
		return Pokemon{}, err
	}

	RespPokemonAbility := Pokemon{}
	err = json.Unmarshal(data, &RespPokemonAbility)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return RespPokemonAbility, nil
}
