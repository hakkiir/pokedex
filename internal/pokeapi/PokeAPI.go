package pokeapi

import (
	"encoding/json"
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
func CommandMap(url string) (locationAreas, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return locationAreas{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	defer res.Body.Close()

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return l, err
	}
	return l, nil
}

func CommandMapb(url *string) (locationAreas, error) { //(next string, previous *string, err error) {

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
		return locationAreas{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	defer res.Body.Close()

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return l, err
	}

	return l, nil

}
