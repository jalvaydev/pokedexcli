package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jalvaydev/pokedexcli/internal/pokecache"
)

func (c *Client) ListPokemon(pageURL *string, cache *pokecache.Cache) (PokemonsResp, error) {
	endpoint := "/pokemon"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return PokemonsResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonsResp{}, err
		}

		if resp.StatusCode > 399 {
			return PokemonsResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonsResp{}, err
		}
		cache.Add(fullURL, data)
	}

	pokemonsResp := PokemonsResp{}
	err := json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return PokemonsResp{}, err
	}

	return pokemonsResp, nil
}

func (c *Client) GetPokemon(pokemonName string, cache *pokecache.Cache) (PokemonResp, error) {
	endpoint := "/pokemon"
	fullURL := baseURL + endpoint + fmt.Sprintf("/%v", pokemonName)

	data, ok := cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return PokemonResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonResp{}, err
		}

		if resp.StatusCode > 399 {
			return PokemonResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonResp{}, err
		}
		cache.Add(fullURL, data)
	}

	pokemonResp := PokemonResp{}
	err := json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil
}
