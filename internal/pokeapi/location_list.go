package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageUrl *string)(locationAreaResponse, error){
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}
	var locationResp locationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return locationAreaResponse{}, err
	}

	return locationResp, nil 
}