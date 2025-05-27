package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	var response RespShallowLocations
	if data, ok := apiCache.Get(url); ok {
		fmt.Println(data)
		if err := json.Unmarshal(data, &response); err != nil {
			return response, err
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	apiCache.Add(url, body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
