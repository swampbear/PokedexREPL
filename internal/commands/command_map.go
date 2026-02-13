package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchMaps(url string, conf *Config) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error, failed to get response: %w", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %w", err)
	}
	locations, err := parseLocationsFromBytes(dat)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	conf.PokeCache.Add(url, dat)

	conf.Next = locations.Next
	conf.Previous = locations.Previous
	if conf.Previous == "" {
		fmt.Println("you're on the first page")
	}

	printLocations(locations)
	return nil

}

func parseLocationsFromBytes(dat []byte) (Locations, error) {
	locations := Locations{}
	if err := json.Unmarshal(dat, &locations); err != nil {
		return Locations{}, fmt.Errorf("Failed to unmarshal body %w", err)
	}
	return locations, nil
}

func CommandMap(conf *Config) error {
	url := conf.Next
	pokeCache, ok := conf.PokeCache.Get(url)
	if ok {
		locations, err := parseLocationsFromBytes(pokeCache)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		conf.Next = locations.Next
		conf.Previous = locations.Previous

		printLocations(locations)

	} else {
		err := fetchMaps(url, conf)
		if err != nil {
			return fmt.Errorf("Error fetching next maps %w", err)
		}

	}
	return nil
}

func CommandBMap(conf *Config) error {
	url := conf.Previous
	if url == "" {
		fmt.Println("You are at page 1")
		return nil
	}

	pokeCache, ok := conf.PokeCache.Get(url)
	if ok {
		locations, err := parseLocationsFromBytes(pokeCache)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		conf.Next = locations.Next
		conf.Previous = locations.Previous
		printLocations(locations)

	} else {
		err := fetchMaps(url, conf)
		if err != nil {
			return fmt.Errorf("Error fetching previous maps %w", err)
		}
	}
	return nil

}

func printLocations(locations Locations) {
	for _, pkmaps := range locations.Results {
		fmt.Println(pkmaps.Name)
	}

}
