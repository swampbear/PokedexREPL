package commands

import "fmt"

func CommandPokedex(conf *Config) error {

	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
