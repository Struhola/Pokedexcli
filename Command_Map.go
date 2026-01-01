package main

import (
	"errors"
	"fmt"
	"path"
)

func commandMapf(cfg *cliConfig, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.List_Locations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	fmt.Printf("ID   | Name\n")
	for _, loc := range locationsResp.Results {
		id := path.Base(loc.URL)
		fmt.Printf("%-4v | %s\n", id, loc.Name)

	}
	return nil
}

func commandMapb(cfg *cliConfig, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.List_Locations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous
	fmt.Printf("ID   | Name\n")
	for _, loc := range locationResp.Results {
		id := path.Base(loc.URL)
		fmt.Printf("%-4v| %s\n", id, loc.Name)
	}
	return nil
}
