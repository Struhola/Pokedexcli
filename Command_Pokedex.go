package main

import (
	"fmt"
)

func commandPokedex(cfg *cliConfig, args ...string) error {
	if len(cfg.pokedex.byName) == 0 {
		return fmt.Errorf("You have no Pokemon in your Pokedex")
	}
	fmt.Printf("Your Pokedex:\n")
	fmt.Printf("ID   | Name\n")
	for _, pokemon := range cfg.pokedex.byName {
		fmt.Printf("%-4v | %s\n", pokemon.ID, pokemon.Name)
	}

	return nil
}
