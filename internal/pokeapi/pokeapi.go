package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cache "github.com/hakkiir/pokedex/internal/pokecache"
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
func CommandMap(url string, c *cache.Cache) (locationAreas, error) {

	//if values are cached, return cached values
	if val, ok := c.Get(url); ok {
		l := locationAreas{}
		err := json.Unmarshal(val, &l)
		//	if err != nil {
		//		return l, err
		//	}
		fmt.Println()
		fmt.Println("returned values from Cache!")
		fmt.Println()
		return l, err
	}

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

	//add results to cache
	c.Add(url, body)

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return l, err
	}
	fmt.Println()
	fmt.Println("returned values from API!")
	fmt.Println()
	return l, nil
}

func CommandMapb(url *string, c *cache.Cache) (locationAreas, error) { //(next string, previous *string, err error) {

	//if values are cached, return cached values
	if val, ok := c.Get(*url); ok {
		l := locationAreas{}
		err := json.Unmarshal(val, &l)
		//if err != nil {
		//	return l, err
		//}
		fmt.Println()
		fmt.Println("returned values from Cache!")
		fmt.Println()
		return l, err
	}

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

	//add results to cache
	c.Add(*url, body)

	l := locationAreas{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return l, err
	}
	fmt.Println()
	fmt.Println("returned values from API!")
	fmt.Println()
	return l, nil

}
