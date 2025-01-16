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
	next     string
	previous *string
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

func commandMap(cfg *config) error {

	//HTTP request and JSON parsing
	//Currently pokeapi.CommandMap also prints the results.. should probably edit to just return values and do printing here
	l, err := pokeapi.CommandMap(cfg.next)

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

func commandMapb(cfg *config) error {

	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return fmt.Errorf("previous page = nil")
	}

	l, err := pokeapi.CommandMapb(cfg.previous)

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
