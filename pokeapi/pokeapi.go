package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type config struct {
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
}

func GetLocations() {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(string(body))
	var locationResponse config
	if err := json.Unmarshal(body, &locationResponse); err != nil {
		log.Fatal(err)
	}
	fmt.Println(locationResponse)

	for _, name := range locationResponse.Results {
		fmt.Printf("%v\n", name)
	}
}
