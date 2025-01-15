package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
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

		commands := getCommands()

		if cb, ok := commands[fw]; ok {
			cb.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
