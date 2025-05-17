package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if entry, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		if err := json.Unmarshal(entry, &locationsResp); err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	if err = json.Unmarshal(dat, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

// GetLocation -
func (c *Client) GetLocation(name string) (RespShallowExplore, error) {
	url := baseURL + "/location-area/" + name

	if entry, ok := c.cache.Get(url); ok {
		exploreResp := RespShallowExplore{}
		if err := json.Unmarshal(entry, &exploreResp); err != nil {
			return RespShallowExplore{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowExplore{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RespShallowExplore{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowExplore{}, err
	}

	c.cache.Add(url, dat)

	exploreResp := RespShallowExplore{}
	if err = json.Unmarshal(dat, &exploreResp); err != nil {
		return RespShallowExplore{}, err
	}

	return exploreResp, nil
}

// GetPokemon -
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if entry, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		if err := json.Unmarshal(entry, &pokemonResp); err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	pokemonResp := Pokemon{}
	if err = json.Unmarshal(dat, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
