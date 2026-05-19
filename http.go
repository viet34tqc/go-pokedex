package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func getIssueData() ([]Issue, error) {
	res, err := http.Get("https://api.boot.dev/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return issues, nil
}
