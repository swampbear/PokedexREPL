package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var c *config

func startREPL() {
	// starts scanner and application loop for user interaction

	c = &config{Next: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		cmd_txt := words[0]
		cmd, exists := getCommands()[cmd_txt]
		if exists {
			err := cmd.callback(c)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("unknown command")
		}

	}

}

func cleanInput(text string) []string {
	new_text := strings.ToLower(text)
	// Fields separates strings into a slice of strings separated by spaces
	words := strings.Fields(new_text)
	return words
}

// struct for standarising cli commands
type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	Next     string
	Previous string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists out pokemon cities",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Lists previous citites",
			callback:    commandBMap,
		},
	}

}
