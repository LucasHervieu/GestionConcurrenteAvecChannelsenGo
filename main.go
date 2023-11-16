// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"ESTIAM/dictionary"
)

func main() {
	d := dictionary.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Add a word and its definition")
		fmt.Println("2. Define a word")
		fmt.Println("3. Remove a word")
		fmt.Println("4. List all words and definitions")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			actionAdd(d, reader)
		case 2:
			actionDefine(d, reader)
		case 3:
			actionRemove(d, reader)
		case 4:
			actionList(d)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1] // Remove newline character

	fmt.Print("Enter the definition: ")
	definition, _ := reader.ReadString('\n')
	definition = definition[:len(definition)-1] // Remove newline character

	d.Add(word, definition)
	fmt.Println("Word added successfully.")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to define: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1] // Remove newline character

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Definition of %s: %s\n", word, entry.Definition)
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1] // Remove newline character

	d.Remove(word)
	fmt.Println("Word removed successfully.")
}

func actionList(d *dictionary.Dictionary) {
	words, entries := d.List()

	if len(words) == 0 {
		fmt.Println("The dictionary is empty.")
		return
	}

	fmt.Println("Words and Definitions:")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entries[word].Definition)
	}
}
