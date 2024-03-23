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

type PokemonResults struct {
	PokemonEncounters []PokemonDetails `json:"pokemon_encounters"`
}

type PokemonDetails struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

var locationResponse config
var pokemonResponse PokemonResults
var pokeApiURL string = "https://pokeapi.co/api/v2/location/"
var pokeApiSearchLocationURL string = "https://pokeapi.co/api/v2/location-area/"
var cache = pokecache.NewCache(50 * time.Second)

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

	body := makeHttpRequest(pokeApiURL)
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

func ExplorePokemons(location string) {
	locationSearchURL := pokeApiSearchLocationURL + location

	if val, hasCache := cache.Get(locationSearchURL); hasCache {
		printPokemons(val)
		return
	}

	body := makeHttpRequest(locationSearchURL)
	cache.Add(locationSearchURL, body)
	printPokemons(body)
}

func printPokemons(body []byte) {
	if err := json.Unmarshal(body, &pokemonResponse); err != nil {
		fmt.Printf("Error : Something went wrong. Please try again with a different location\n")
		return
	}

	for _, p := range pokemonResponse.PokemonEncounters {
		fmt.Printf("%v \n", p.Pokemon.Name)
	}
}

func makeHttpRequest(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	return body
}
