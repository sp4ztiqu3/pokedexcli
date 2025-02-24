package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{}

	commandExit := func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
	commandHelp := func() error {
		fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
		for _, c := range commands {
			fmt.Printf("%s: %s\n", c.name, c.description)
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

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		found := false
		input := cleanInput(scanner.Text())
		for _, c := range commands {
			if c.name == input[0] {
				c.callback()
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
