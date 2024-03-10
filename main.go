package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var cliName string = "pokedex"

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
	printPrompt()
	for reader.Scan() {
		text := reader.Text()
		if command, ok := commands[text]; ok {
			command.callback()
		} else {

		}
		printPrompt()
	}
	fmt.Println()
}

func printPrompt() {
	fmt.Print(cliName, "> ")
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
