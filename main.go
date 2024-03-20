package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CoderParth/pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand) error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Displays the names of next 20 location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous 20 locations",
		callback:    commandMapB,
	},
}

func commandHelp(cMap map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for cKey, v := range cMap {
		fmt.Printf("%v: %v\n", cKey, v.description)
	}
	return nil
}

func commandExit(cMap map[string]cliCommand) error {
	os.Exit(0)
	return nil
}

func commandMap(cMap map[string]cliCommand) error {
	pokeapi.GetLocations()
	return nil
}

func commandMapB(cMap map[string]cliCommand) error {
	return nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println()
		command, ok := commands[text]
		if ok {
			command.callback(commands)
		} else {
			fmt.Println("Command not found")
		}
	}
}
