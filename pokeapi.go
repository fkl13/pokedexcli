package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
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

func callLocationAreaByName(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}
		cache.Add(url, body)
	}

	locations := Location{}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return Location{}, err
	}

	return locations, nil
}

type Location struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(url, body)

	return pokemon, nil
}

type Pokemon struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	IsDefault              bool   `json:"is_default"`
	Order                  int    `json:"order"`
	Weight                 int    `json:"weight"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
