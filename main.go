package main

import (
	"bufio"
	"fmt"
	"os"
	"ESTIAM/dictionary"
	"strings"
)

func main() {
	d := dictionary.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n----- Dictionary Menu -----")
		fmt.Println("1. Add a word and its definition")
		fmt.Println("2. Define a word")
		fmt.Println("3. Remove a word")
		fmt.Println("4. List all words and definitions")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice (1-5): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			actionAdd(d, reader)
		case "2":
			actionDefine(d, reader)
		case "3":
			actionRemove(d, reader)
		case "4":
			actionList(d)
		case "5":
			fmt.Println("Exiting the program. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Enter the definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	d.Add(word, definition)
	fmt.Println("Word added successfully.")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to define: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Definition of '%s': %s\n", word, entry.Definition)
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	d.Remove(word)
	fmt.Println("Word removed successfully.")
}

func actionList(d *dictionary.Dictionary) {
	words, entries := d.List()

	if len(words) == 0 {
		fmt.Println("The dictionary is empty.")
		return
	}

	fmt.Println("\nWords and Definitions:")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entries[word].Definition)
	}
}
