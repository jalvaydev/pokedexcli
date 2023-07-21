package main

import (
	"time"

	"github.com/jalvaydev/pokedexcli/internal/pokeapi"
	"github.com/jalvaydev/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	pokeCache           *pokecache.Cache
	caughtPokemon       map[string]pokeapi.PokemonResp
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		caughtPokemon: map[string]pokeapi.PokemonResp{},
		pokeapiClient: pokeapi.NewClient(),
		pokeCache:     pokecache.NewCache(5 * time.Second),
	}

	startRepl(&cfg)
}
