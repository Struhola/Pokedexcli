package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func cleanInput(text string) []string {
// 	var result []string
// 	split_text := strings.Fields(strings.ToLower(text))
// 	fmt.Println(split_text)
// 	for _, word := range split_text {
// 		result = append(result, strings.TrimSpace(word))
// 	}
// 	return result
// }

func main() {
	ui := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ui.Scan() {
			clear_ui := strings.ToLower(strings.TrimSpace(ui.Text()))
			words := strings.Fields(clear_ui)
			if len(words) > 0 {
				command := words[0]
				fmt.Printf("Your command was: %s\n", command)
			}

		}

	}
}
