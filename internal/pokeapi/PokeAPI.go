package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationAreas struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

/*
	type Config struct {
		Next     string
		Previous *string
	}
*/
func CommandMap(url string) (next string, previous *string, err error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	defer res.Body.Close()

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return "", nil, err
	}

	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}

	return l.Next, l.Previous, nil
}

func CommandMapb(url *string) (next string, previous *string, err error) {

	if url == nil {
		fmt.Println("you're on the first page")
		return "", nil, nil
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	defer res.Body.Close()

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return "", nil, err
	}

	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}

	return l.Next, l.Previous, nil
}
