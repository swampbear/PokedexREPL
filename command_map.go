package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MapGroup struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchMaps(url string, conf *config) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error, failed to get response: %w", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %w", err)
	}
	maps := MapGroup{}
	if err = json.Unmarshal(dat, &maps); err != nil {
		return fmt.Errorf("Failed to unmarshal body %w", err)
	}

	conf.Next = maps.Next
	conf.Previous = maps.Previous
	if conf.Previous == "" {
		fmt.Println("you're on the first page")
	}

	for _, pkmaps := range maps.Results {
		fmt.Println(pkmaps.Name)
	}
	return nil

}

func commandMap(conf *config) error {
	url := conf.Next
	err := fetchMaps(url, conf)
	if err != nil {
		return fmt.Errorf("Error fetching next maps %w", err)
	}
	return nil
}

func commandBMap(conf *config) error {
	url := conf.Previous
	err := fetchMaps(url, conf)
	if err != nil {
		return fmt.Errorf("Error fetching previous maps %w", err)
	}
	return nil

}
