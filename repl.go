package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/elifortmort/pokedex/internal/pokecache"
)

func replLoop(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := c.commands

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() && len(scanner.Text()) > 0 {
			text := scanner.Text()
			wrds := cleanInput(text)
			if len(wrds) >= 1 {
				_, ok := cmd[wrds[0]]
				if ok {
					if len(wrds) > 1 {
						cmd[wrds[0]].callback(c, wrds[1])
					} else {
						cmd[wrds[0]].callback(c, "")
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		} else {
			fmt.Println("Command required...")
		}
	}
}

func cleanInput(text string) []string {
	lwr := strings.ToLower(text)
	trim := strings.TrimSpace(lwr)
	words := strings.Fields(trim)
	return words
}

type config struct {
	commands map[string]cliCommand
	nextUrl string
	previousUrl string
	cache *pokecache.Cache
	pokedex map[string]Pokemon
}

type Pokemon struct {
	height int
	weight int
	stats map[string]int
	types []string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 areas and any further calls to map will display the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous list of 20 areas if map was called more than once",
			callback:    commandMapb,
		},
		"explore": {
			name:		 "explore <area>",
			description: "Displays a list of possible pokemon to encounter in a specific area",
			callback: commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon>",
			description: "Attempts a capture of specified pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name:		 "inspect <pokemon>",
			description: "Attempts inspection of specified pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Display a list of caught pokemon",
			callback: commandPokedex,
		},
	}
}
