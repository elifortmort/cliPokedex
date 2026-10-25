package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"math/rand"
)

func commandCatch(c *config, str string) error {

	type pStat struct {
		Name string `json:"name"`
	}

	type PokemonStat struct {
		Base int `json:"base_stat"`
		Info pStat `json:"stat"` 
	}

	type pType struct {
		Name string `json:"name"`
	}

	type PokemonType struct {
		Type pType `json:"type"`
	}

	type pokeByte struct {
		Name string `json:"name"`
		Xp int `json:"base_experience"`
		Height int `json:"height"`
		Weight int `json:"weight"`
		Stats []PokemonStat `json:"stats"`
		Types []PokemonType `json:"types"`
	}

	var url string
	chance := 45


	if len(str) <= 0 {
		fmt.Println("Pokemon name needed to attempt catch...")
		return nil
	}
	url = "https://pokeapi.co/api/v2/pokemon/" + str

	// Look to cache if an entry exists
	if val, ok := c.cache.Get(url); ok {
		pokeB := pokeByte{}
		errer := json.Unmarshal(val, &pokeB)
		if errer != nil {
			fmt.Println(errer)
		}

		//Create pokemon
		pokemon := Pokemon{
			height: pokeB.Height,
			weight: pokeB.Weight,
			stats: make(map[string]int),
			types: []string{},
		}

		//Populate stats and types
		for _, stat := range pokeB.Stats {
			pokemon.stats[stat.Info.Name] = stat.Base
		}
		for _, t := range pokeB.Types {
			pokemon.types = append(pokemon.types, t.Type.Name)
		}

		//Decide wether pokemon was caught and added to the pokedex or not, based on pokemons base xp stat
		fmt.Printf("Throwing a Pokeball at %s...\n", str)
		if chance > rand.Intn(pokeB.Xp) {
			fmt.Printf("%s was caught!\n", str)
			c.pokedex[str] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", str)
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
		fmt.Printf("Response failed with status code: %d and\nbody: %s\nCheck if the pokemon is real with the explore command...\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Add a cache entry
	c.cache.Add(url, body)

	// Unmarshal the byte
	pokeB := pokeByte{}
	errer := json.Unmarshal(body, &pokeB)
	if errer != nil {
		fmt.Println(errer)
	}

	//Create pokemon
	pokemon := Pokemon{
		height: pokeB.Height,
		weight: pokeB.Weight,
		stats: make(map[string]int),
		types: []string{},
	}

	//Populate stats and types
	for _, stat := range pokeB.Stats {
		pokemon.stats[stat.Info.Name] = stat.Base
	}
	for _, t := range pokeB.Types {
		pokemon.types = append(pokemon.types, t.Type.Name)
	}

	//Decide wether pokemon was caught and added to the pokedex or not, based on pokemons base xp stat
	fmt.Printf("Throwing a Pokeball at %s...\n", str)
	if chance > rand.Intn(pokeB.Xp) {
		fmt.Printf("%s was caught!\n", str)
		c.pokedex[str] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", str)
	}

	return nil
}