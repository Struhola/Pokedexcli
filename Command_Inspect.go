package main

import (
	"Pokedexcli/internal/pokeapi"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func commandInspect(cfg *cliConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon name or ID")
	}

	Key := args[0]
	var msg strings.Builder
	var data *pokeapi.Pokemon
	var ok bool

	if data, ok = cfg.pokedex.GetByName(Key); !ok {
		KeyInt, err := strconv.Atoi(Key)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if data, ok = cfg.pokedex.GetByID(KeyInt); !ok {
			return fmt.Errorf("you have not caught this Pokemon\n")
		}
	}

	fmt.Fprintf(&msg, "Name: %v\n", data.Name)
	fmt.Fprintf(&msg, "Height: %v\n", data.Height)
	fmt.Fprintf(&msg, "Weight: %v\n", data.Weight)
	fmt.Fprintf(&msg, "Stats:\n")
	for _, st := range data.Stats {
		fmt.Fprintf(&msg, "  -%s: %v\n", st.Stat.Name, st.BaseStat)
	}
	fmt.Fprintf(&msg, "Types:\n")
	for _, tp := range data.Types {
		fmt.Fprintf(&msg, "  -%v\n", tp.Type.Name)
	}

	fmt.Println(msg.String())
	return nil
}
