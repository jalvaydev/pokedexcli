package main

import (
	"errors"
	"fmt"
)

func commandExplore(params []string, cfg *config) error {
	if len(params) == 0 {
		return errors.New("location paramater required")
	}

	locationAreaName := params[0]

	fmt.Printf("Exploring %v...\n", locationAreaName)

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName, cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range resp.PokemonEncounters {
		fmt.Printf(" - %v\n", v.Pokemon.Name)
	}

	return nil
}
