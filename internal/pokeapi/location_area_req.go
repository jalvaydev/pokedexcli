package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jalvaydev/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreasResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreasResp{}, err
		}

		if resp.StatusCode > 399 {
			return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreasResp{}, err
		}
		cache.Add(fullURL, data)
	}

	locationAreasResp := LocationAreasResp{}
	err := json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationName string, cache *pokecache.Cache) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + fmt.Sprintf("/%v", locationName)

	data, ok := cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreaResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResp{}, err
		}

		if resp.StatusCode > 399 {
			return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResp{}, err
		}
		cache.Add(fullURL, data)
	}

	locationAreaResp := LocationAreaResp{}
	err := json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreaResp, nil
}
