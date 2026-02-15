package commands

import (
	"github.com/swampbear/pokedexcli/internal/models"
	"github.com/swampbear/pokedexcli/internal/pokecache"
)

type Config struct {
	Pokedex   map[string]models.Pokemon
	Action    string
	PokeCache pokecache.Cache
	Next      string
	Previous  string
}

// CliCommand is a struct for standardizing cli commands
type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Lists out pokemon cities",
			Callback:    CommandMap,
		},
		"bmap": {
			Name:        "bmap",
			Description: "Lists previous citites",
			Callback:    CommandBMap,
		},
		"explore": {
			Name:        "explore",
			Description: "Lists previous cities",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Command for catching pokemon given as second argument",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect details of a given pokemon in your pokedex",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists out all caught pokemon",
			Callback:    CommandPokedex,
		},
	}
}
