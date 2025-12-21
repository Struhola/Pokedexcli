package main

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
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
			description: "Displays location results next page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays location results previous page",
			callback:    commandMapb,
		},
	}
}
