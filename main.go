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
	callback    func() error
}

var commands map[string]cliCommand = map[string]cliCommand{
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
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		reader.Scan()
		text := reader.Text()
		commandName := cleanInput(text)[0]
		if command, ok := commands[commandName]; ok {
			err := command.callback()
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

func commandHelp() error {
	message := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Println(message)
	return nil
}

func commandExit() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	os.Exit(0)
	return nil
}
