package commands

type Config struct {
	Next     string
	Previous string
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
	}
}
