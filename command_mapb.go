package main

import (
	"errors"
	"fmt"
)

func commandMapb(params []string, cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("no previous page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL, cfg.pokeCache)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
