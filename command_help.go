package main

import (
	"fmt"
)

func commandHelp(c *config, str string) error {
	fmt.Println("Welcome to the Pokedex!\n")
	cmd := c.commands
	for _, item := range cmd {
		fmt.Printf("%s:\n	%s\n", item.name, item.description)
		fmt.Println("")
	}
	return nil
}