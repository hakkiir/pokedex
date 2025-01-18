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

	c := pcache.NewCache(time.Minute * 5)

	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	cfg := config{
		next:     &url,
		previous: nil,
		cache:    c,
	}
	for {
		param := ""
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
		if len(input) > 1 {
			param = input[1]
		}

		commands := getCommands(&cfg, "")

		if cb, ok := commands[fw]; ok {
			cb.callback(&cfg, param)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
