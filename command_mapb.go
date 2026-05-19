package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(config *ParamsConfig) error {
	if config.Prev == "" {
		return fmt.Errorf("you're on the first page")
	}

	res, err := http.Get(config.Prev)
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
