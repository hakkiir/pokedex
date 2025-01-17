package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pcache "github.com/hakkiir/pokedex/internal/pokecache"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	c := pcache.NewCache(time.Second * 20)

	cfg := config{
		next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		previous: nil,
		cache:    c,
	}
	for {
		fmt.Print("Pokedex > ")
		err := scanner.Scan()
		if !err {
			fmt.Println(err)
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		fw := input[0]

		commands := getCommands(&cfg)

		if cb, ok := commands[fw]; ok {
			cb.callback(&cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
