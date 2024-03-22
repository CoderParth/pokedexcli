package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/CoderParth/pokedexcli/pokecache"
)

type config struct {
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
}

var locationResponse config
var pokeApiURL string = "https://pokeapi.co/api/v2/location/"
var cache = pokecache.NewCache(5 * time.Second)

func GetNextLocations() {
	if locationResponse.Next != "" {
		pokeApiURL = locationResponse.Next
	}

	getLocations()
}

func GetPrevLocations() {
	if locationResponse.Previous != "" {
		pokeApiURL = locationResponse.Previous
		getLocations()
		return
	}

	fmt.Println("Error: You do not have any previously searched locations.")
}

func getLocations() {
	if val, hasCache := cache.Get(pokeApiURL); hasCache {
		printLocations(val)
		return
	}

	res, err := http.Get(pokeApiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}

	cache.Add(pokeApiURL, body)
	printLocations(body)
}

func printLocations(body []byte) {
	if err := json.Unmarshal(body, &locationResponse); err != nil {
		log.Fatal(err)
	}

	for _, name := range locationResponse.Results {
		fmt.Printf("%v\n", name.Name)
	}
}
