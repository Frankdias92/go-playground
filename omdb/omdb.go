package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
}
type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey, title string) (Result, error) {
	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s", apiKey, title)

	resp, err := http.Get(url)

	if err != nil {
		return Result{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}

	defer resp.Body.Close()

	var result Result

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("failed to decode response from omdb: %w", err)
	}

	return result, nil
}
