package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(params []string, cfg *config) error {
	if len(params) == 0 {
		return errors.New("pokemon name is required")
	}

	pokemonName := params[0]

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName, cfg.pokeCache)
	if err != nil {
		return err
	}

	pokemonCatchRate := resp.BaseExperience / 32 * 5
	randomChance := rand.Intn(100)
	if randomChance >= pokemonCatchRate {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = resp
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}
