package main

import (
	"fmt"
)

func commandInspect(c *config, str string) error {
	
	if _, ok := c.pokedex[str]; ok {
		pokemon := c.pokedex[str]
		stats := pokemon.stats
		types := pokemon.types
		fmt.Printf(
			"Name: %s\nHeight: %d\nWeight: %d\nStats:\n", 
			str, 
			pokemon.height, 
			pokemon.weight,
		)
		for i, stat := range stats {
			fmt.Printf(" -%s: %d\n", i, stat)
		}
		fmt.Println("Types:")
		for _, t := range types {
			fmt.Printf(" - %s\n", t)
		}
	} else {
		fmt.Println("The specified pokemon is either not in the pokedex or does not exist...")
		fmt.Println("Try the catch command to attempt to add the pokemon to the pokedex!")
	}

	return nil
}