package main

import (
	"fmt"
	"internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next     string
	Previous *string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {

	commands := getCommands(cfg)

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}

func getCommands(cfg *config) map[string]cliCommand {

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
	}
	return commands
}

func commandMap(cfg *config) error {
	next, previous, err := pokeapi.CommandMap(cfg.Next)
	if err != nil {
		fmt.Println(err)
	}
	cfg.Next = next
	cfg.Previous = previous
	return nil
}

func commandMapb(cfg *config) error {

	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return fmt.Errorf("previous page = nil")
	}
	next, previous, err := pokeapi.CommandMapb(cfg.Previous)
	if err != nil {
		fmt.Println(err)
	}
	cfg.Next = next
	cfg.Previous = previous
	return nil
}
