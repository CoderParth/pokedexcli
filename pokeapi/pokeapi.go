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
	}

	fmt.Println("Error: You do not have any previously searched locations.")
}

func getLocations() {
	res, err := http.Get(pokeApiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	// fmt.Println(string(body))
	if err := json.Unmarshal(body, &locationResponse); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(locationResponse)

	for _, name := range locationResponse.Results {
		fmt.Printf("%v\n", name.Name)
	}
}
