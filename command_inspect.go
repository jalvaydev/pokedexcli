package main

import (
	"errors"
	"fmt"
)

func commandInspect(params []string, cfg *config) error {
	if len(params) == 0 {
		return errors.New("pokemon name is required")
	}

	pokemonName := params[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf(" -%v\n", v.Type.Name)
	}
	return nil
}
