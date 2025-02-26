package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	nextUrl string
	prevUrl string
}

type pokeLocationAreas struct {
	Count   int     `json:"count"`
	NextUrl *string `json:"next"`
	PrevUrl *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{}

	commandExit := func(conf *config) error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
	commandHelp := func(conf *config) error {
		fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
		for _, c := range commands {
			fmt.Printf("%s: %s\n", c.name, c.description)
		}
		return nil
	}
	commandMap := func(conf *config) error {
		if conf.nextUrl == "" {
			fmt.Println("no next page available")
			return nil
		}
		res, err := http.Get(conf.nextUrl)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		defer res.Body.Close()

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		var locationAreas pokeLocationAreas
		if err = json.Unmarshal(bodyData, &locationAreas); err != nil {
			return fmt.Errorf("error unmarshalling json: %w", err)
		}

		if locationAreas.NextUrl != nil {
			conf.nextUrl = *locationAreas.NextUrl
		} else {
			conf.nextUrl = ""
		}
		if locationAreas.PrevUrl != nil {
			conf.prevUrl = *locationAreas.PrevUrl
		} else {
			conf.prevUrl = ""
		}

		for _, loc := range locationAreas.Results {
			fmt.Println(loc.Name)
		}

		return nil
	}
	commandPMap := func(conf *config) error {
		if conf.prevUrl == "" {
			fmt.Println("no previous page available")
			return nil
		}
		res, err := http.Get(conf.prevUrl)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		defer res.Body.Close()

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		var locationAreas pokeLocationAreas
		if err = json.Unmarshal(bodyData, &locationAreas); err != nil {
			return fmt.Errorf("error unmarshalling json: %w", err)
		}

		if locationAreas.NextUrl != nil {
			conf.nextUrl = *locationAreas.NextUrl
		} else {
			conf.nextUrl = ""
		}
		if locationAreas.PrevUrl != nil {
			conf.prevUrl = *locationAreas.PrevUrl
		} else {
			conf.prevUrl = ""
		}

		for _, loc := range locationAreas.Results {
			fmt.Println(loc.Name)
		}

		return nil
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Get the next 20 locations from PokeAPI",
		callback:    commandMap,
	}
	commands["pmap"] = cliCommand{
		name:        "pmap",
		description: "Get the previous 20 locations from PokeAPI",
		callback:    commandPMap,
	}

	conf := config{
		nextUrl: "https://pokeapi.co/api/v2/location-area",
		prevUrl: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		found := false
		input := cleanInput(scanner.Text())
		for _, c := range commands {
			if c.name == input[0] {
				c.callback(&conf)
				found = true
			}
		}
		if !found {
			fmt.Printf("Unknown command: %s\n", input[0])
		}
	}
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}
