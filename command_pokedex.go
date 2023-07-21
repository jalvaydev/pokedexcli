package main

import (
	"errors"
	"fmt"
)

func commandPokedex(params []string, cfg *config) error {
	fmt.Printf("Your Pokedex:\n")
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("pokedex is empty")
	}
	for _, v := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", v.Name)
	}
	return nil
}
