package main

import (
	"Pokedexcli/internal/pokeapi"
	"Pokedexcli/internal/pokecache"
)

type cliConfig struct {
	pokeapiClient    pokeapi.Client
	cache            pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          *Pokedex
}
