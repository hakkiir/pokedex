package main

import (
	"fmt"
	"os"

	"github.com/hakkiir/pokedex/internal/pokeapi"
	pcache "github.com/hakkiir/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, param string) error
}

type config struct {
	next     *string
	previous *string
	cache    pcache.Cache
}

func getCommands(cfg *config, param string) map[string]cliCommand {

	commands := map[string]cliCommand{
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
			description: "print next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "print previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore area",
			callback:    commandExplore,
		},
	}
	return commands
}

func commandExit(cfg *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param string) error {

	commands := getCommands(cfg, "")

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}

func commandMap(cfg *config, param string) error {

	l, err := pokeapi.GetLocations(cfg.next, cfg.cache)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = l.Next
	cfg.previous = l.Previous
	return nil
}

func commandMapb(cfg *config, param string) error {

	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return fmt.Errorf("previous page = nil")
	}

	l, err := pokeapi.GetLocations(cfg.previous, cfg.cache)

	if err != nil {
		fmt.Println(err)
	}

	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = l.Next
	cfg.previous = l.Previous
	return nil
}

func commandExplore(cfg *config, param string) error {

	fmt.Printf("Exploring %s...\n", param)
	url := "https://pokeapi.co/api/v2/location-area/" + param
	e, err := pokeapi.Explore(url, cfg.cache)

	if err != nil {
		fmt.Println("No pokemons found!")
		fmt.Println(err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range e.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}
