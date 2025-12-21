package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliConfig struct {
	PreviousURL *string
	NextURL     *string
}

func commandExit(cfg *cliConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *cliConfig) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *cliConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextURL != nil {
		url = *cfg.NextURL
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	response := locationAreaRespone{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	cfg.PreviousURL = response.Previous
	cfg.NextURL = response.Next

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *cliConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.PreviousURL != nil {
		url = *cfg.PreviousURL
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	response := locationAreaRespone{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	cfg.NextURL = cfg.PreviousURL
	cfg.PreviousURL = response.Previous

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	return nil
}
