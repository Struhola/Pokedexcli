package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *cliConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon name or ID")
	}

	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.Pokemon_Information(pokemonName)
	if err != nil {
		return err
	}

	catchDifficulty := 500
	catchRoll := rand.Intn(catchDifficulty)
	pokemonBaseExp := pokemonResp.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if catchRoll > pokemonBaseExp {
		fmt.Printf("Success! %s has been caught and added to your Pokedex\n", pokemonResp.Name)
		cfg.pokedex.Add(pokemonResp)
	} else {
		fmt.Printf("You have failed to catch %s \n", pokemonResp.Name)
	}
	//Pokemon := fmt.Sprintf("%v. %s", pokemonResp.ID, pokemonResp.Name)
	//fmt.Printf("Exploring %v...\n", Explored_Area)

	return nil
}
