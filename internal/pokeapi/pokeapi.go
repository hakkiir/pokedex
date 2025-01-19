package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cache "github.com/hakkiir/pokedex/internal/pokecache"
)

func GetLocations(cfgUrl *string, c cache.Cache) (locationAreas, error) {

	url := *cfgUrl
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

func Explore(url string, c cache.Cache) (explore, error) {

	e := explore{}
	//if values are cached, return cached values
	if val, ok := c.Get(url); ok {
		err := json.Unmarshal(val, &e)
		//	if err != nil {
		//		return l, err
		//	}
		fmt.Println()
		fmt.Println("returned values from Cache!")
		fmt.Println()
		return e, err
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return e, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return e, err
	}

	defer res.Body.Close()

	//add results to cache
	c.Add(url, body)

	err = json.Unmarshal(body, &e)
	if err != nil {
		return e, err
	}
	fmt.Println()
	fmt.Println("returned values from API!")
	fmt.Println()
	return e, nil
}

func Catch(url string, c cache.Cache) (Pokemon, error) {

	p := Pokemon{}

	//if values are cached, return cached values
	if val, ok := c.Get(url); ok {
		err := json.Unmarshal(val, &p)
		fmt.Println()
		fmt.Println("returned values from Cache!")
		fmt.Println()
		return p, err
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return p, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	defer res.Body.Close()

	//add results to cache
	c.Add(url, body)

	err = json.Unmarshal(body, &p)
	if err != nil {
		return p, err
	}
	fmt.Println()
	fmt.Println("returned values from API!")
	fmt.Println()
	return p, nil
}
