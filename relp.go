package main

import (
	"fmt"
	"go-http-client/internal/pokeapi"
	"strings"

	"github.com/eiannone/keyboard"
)

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type commandHistory struct {
	entries []string
	limit   int
}

func newCommandHistory(limit int) *commandHistory {
	return &commandHistory{
		entries: make([]string, 0, limit),
		limit:   limit,
	}
}

func (h *commandHistory) add(command string) {
	command = strings.TrimSpace(command)
	if command == "" {
		return
	}
	if len(h.entries) > 0 && h.entries[len(h.entries)-1] == command {
		return
	}
	if len(h.entries) == h.limit {
		h.entries = h.entries[1:]
	}
	h.entries = append(h.entries, command)
}

func cleanInput(text string) []string {
	return strings.Fields(text)
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
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your pokedex",
			callback:    commandPokedex,
		},
	}
}

// \r moves cursor to start of line.
// \033[2K wipes the line
// print fresh prompt + input
func redrawInput(prompt string, text []rune) {
	fmt.Printf("\r\033[2K%s%s", prompt, string(text))
}

func startRepl(cfg *Config) {
	const prompt = "Pokedex > "
	history := newCommandHistory(100)
	historyIndex := 0
	input := make([]rune, 0, 128)

	err := keyboard.Open()
	if err != nil {
		fmt.Println("failed to initialize keyboard input:", err)
		return
	}
	defer keyboard.Close()

	for {
		redrawInput(prompt, input)
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("\ninput error:", err)
			return
		}

		switch key {
		case keyboard.KeyCtrlC:
			fmt.Println()
			return
		case keyboard.KeyEnter:
			fmt.Println()
			line := string(input)
			input = input[:0]
			historyIndex = len(history.entries)

			words := cleanInput(line)
			if len(words) == 0 {
				continue
			}

			history.add(line)
			historyIndex = len(history.entries)

			commandName := words[0]
			command, exists := getCommands()[commandName]
			if exists {
				err := command.callback(cfg, words[1:]...)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			fmt.Println("Unknown command")
			continue
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		case keyboard.KeyArrowUp:
			if len(history.entries) == 0 {
				continue
			}
			if historyIndex > 0 {
				historyIndex--
			}
			input = []rune(history.entries[historyIndex])
		case keyboard.KeyArrowDown:
			if len(history.entries) == 0 {
				continue
			}
			if historyIndex < len(history.entries)-1 {
				historyIndex++
				input = []rune(history.entries[historyIndex])
			} else {
				historyIndex = len(history.entries)
				input = input[:0]
			}
		default:
			if key == 0 && char != 0 {
				input = append(input, char)
				historyIndex = len(history.entries)
			}
		}
	}
}
