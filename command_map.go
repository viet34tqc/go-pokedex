package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreasResponse struct {
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(config *ParamsConfig) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var locationAreas LocationAreasResponse
	err = json.NewDecoder(res.Body).Decode(&locationAreas)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	config.Prev = locationAreas.Previous
	config.Next = locationAreas.Next
	return nil
}
