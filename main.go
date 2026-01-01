package main

import (
	"Pokedexcli/internal/pokeapi"
	"Pokedexcli/internal/pokecache"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, pokecache.NewCache(5*time.Second))
	cfg := &cliConfig{
		pokeapiClient: pokeClient,
		pokedex:       NewPokedex(),
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
