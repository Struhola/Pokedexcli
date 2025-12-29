package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *cliConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name or ID")
	}
	areaName := args[0]
	locationResp, err := cfg.pokeapiClient.Location_Deep_Dive(areaName)
	if err != nil {
		return err
	}
	Explored_Area := fmt.Sprintf("%v. %s", locationResp.ID, locationResp.Name)
	fmt.Printf("Exploring %v...\n", Explored_Area)
	fmt.Printf("Found Pokemon:\n")
	for _, Encounter := range locationResp.PokemonEncounters {
		fmt.Println(Encounter.Pokemon.Name)
	}
	return nil
}
