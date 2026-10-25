package main

import (
	"fmt"
)

func commandPokedex(c *config, str string) error {
	
	if len(c.pokedex) <= 0 {
		fmt.Println("The pokedex is empty...")
		fmt.Println("Try catching some pokemon with the catch command!")
		return nil
	} else {
		fmt.Println("Your Pokedex:")
		for i, _ := range c.pokedex {
			fmt.Printf(" - %s\n", i)
		}
	}

	return nil
}