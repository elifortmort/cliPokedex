package main

import (
	"fmt"
	"io"
	//"log"
	"net/http"
	"encoding/json"
)

func commandExplore(c *config, str string) error {

	type pokemon struct {
		Name string `json:"name"`
	}

	type encounter struct {
		Pokemon pokemon `json:"pokemon"`
	}

	type areaByte struct {
		Name string `json:"name"`
		Encounters []encounter `json:"pokemon_encounters"`
	}

	var url string

	if len(str) <= 0 {
		fmt.Println("Area name needed to explore...")
		return nil
	}
	url = "https://pokeapi.co/api/v2/location-area/" + str


	// Look to cache if an entry exists
	if val, ok := c.cache.Get(url); ok {
		areaB := areaByte{}
		errer := json.Unmarshal(val, &areaB)
		if errer != nil {
			fmt.Println(errer)
		}

		for _, encounter := range areaB.Encounters {
			fmt.Printf("%s\n", encounter.Pokemon.Name)
		}
		return nil
	}

	// Call upon data and check if retrieval was succesful
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\nCheck if the place is real with the map command...\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Add a cache entry
	c.cache.Add(url, body)

	// Unmarshal the byte
	areaB := areaByte{}
	errer := json.Unmarshal(body, &areaB)
	if errer != nil {
		fmt.Println(errer)
	}

	for _, encounter := range areaB.Encounters {
		fmt.Printf("%s\n", encounter.Pokemon.Name)
	}

	return nil
}