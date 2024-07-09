package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationsArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func callLocationArea(url string) (LocationsArea, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationsArea{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationsArea{}, err
		}
		cache.Add(url, body)
	}

	locations := LocationsArea{}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return LocationsArea{}, err
	}

	return locations, nil
}
