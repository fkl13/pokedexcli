package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cliName string = "Pokedex"

type cliCommand struct {
	name        string
	description string
	callback    func(p *pagination) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

type pagination struct {
	next     *string
	previous *string
}

func (p *pagination) String() string {
	next := "nil"
	if p.next != nil {
		next = *p.next
	}

	prev := "nil"
	if p.previous != nil {
		prev = *p.previous
	}
	return fmt.Sprintf("Next: %s, Previous: %s", next, prev)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	apiURL := "https://pokeapi.co/api/v2/location-area/"
	pagination := pagination{
		next:     &apiURL,
		previous: nil,
	}
	for {
		printPrompt()
		reader.Scan()

		text := reader.Text()
		input := cleanInput(text)
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		if command, ok := getCommands()[commandName]; ok {
			err := command.callback(&pagination)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	inputWords := strings.Fields(input)
	return inputWords
}

func commandHelp(p *pagination) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()

	return nil
}

func commandExit(p *pagination) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	os.Exit(0)
	return nil
}

func commandMap(p *pagination) error {
	if p.next == nil {
		return fmt.Errorf("already on the last page")
	}

	locations, err := callLocationArea(*p.next)
	if err != nil {
		return err
	}

	p.previous = &locations.Previous
	p.next = &locations.Next

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapb(p *pagination) error {
	if p.previous == nil {
		return fmt.Errorf("already on the first page")
	}

	locations, err := callLocationArea(*p.previous)
	if err != nil {
		return err
	}

	p.previous = &locations.Previous
	p.next = &locations.Next

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}
