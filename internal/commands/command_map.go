package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cities struct {
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
	cities, err := parseCitiesFromBytes(dat)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	conf.PokeCache.Add(url, dat)

	conf.Next = cities.Next
	conf.Previous = cities.Previous
	if conf.Previous == "" {
		fmt.Println("you're on the first page")
	}

	for _, pkmaps := range cities.Results {
		fmt.Println(pkmaps.Name)
	}
	return nil

}

func parseCitiesFromBytes(dat []byte) (Cities, error) {
	cities := Cities{}
	if err := json.Unmarshal(dat, &cities); err != nil {
		return Cities{}, fmt.Errorf("Failed to unmarshal body %w", err)
	}
	return cities, nil
}

func CommandMap(conf *Config) error {
	url := conf.Next
	pokeCache, ok := conf.PokeCache.Get(url)
	if ok {
		cities, err := parseCitiesFromBytes(pokeCache)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		conf.Next = cities.Next
		conf.Previous = cities.Previous

		fmt.Println("THIS WAS CACHES")
		for _, pkmaps := range cities.Results {
			fmt.Println(pkmaps.Name)
		}
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
		cities, err := parseCitiesFromBytes(pokeCache)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		conf.Next = cities.Next
		conf.Previous = cities.Previous

		for _, pkmaps := range cities.Results {
			fmt.Println(pkmaps.Name)
		}
	} else {
		err := fetchMaps(url, conf)
		if err != nil {
			return fmt.Errorf("Error fetching previous maps %w", err)
		}
	}
	return nil

}
