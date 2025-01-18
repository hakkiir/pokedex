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
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type explore struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

/*
	type Config struct {
		Next     string
		Previous *string
	}
*/
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
