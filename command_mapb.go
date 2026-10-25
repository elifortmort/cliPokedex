package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMapb(c *config, str string) error {

	type area struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	type areaByte struct {
		Count int `json:"count"`
		Next string `json:"next"`
		Previous string `json:"previous"`
		Results []area `json:"results"`
	}

	url := c.previousUrl

	if url == "" {
		fmt.Println("No previous list to look up...")
		return nil
	}

	// Look to cache if an entry exists
	if val, ok := c.cache.Get(url); ok {
		areaB := areaByte{}
		errer := json.Unmarshal(val, &areaB)
		if errer != nil {
			fmt.Println(errer)
		}

		c.nextUrl = areaB.Next
		c.previousUrl = areaB.Previous

		for _, area := range areaB.Results {
			fmt.Printf("%s\n", area.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Add a cache entry
	c.cache.Add(url, body)

	areaB := areaByte{}
	errer := json.Unmarshal(body, &areaB)
	if errer != nil {
		fmt.Println(errer)
	}

	c.nextUrl = areaB.Next
	c.previousUrl = areaB.Previous

	fmt.Printf("--------------------new entry----------------------------\n")
	for _, area := range areaB.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}