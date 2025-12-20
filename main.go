package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()
	ui := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ui.Scan()
		words := cleanInput(ui.Text())
		if len(words) == 0 {
			continue
		}
		ui_command := words[0]
		sys_command, exists := commands[ui_command]
		if exists {
			err := sys_command.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
