// dictionary.go

package dictionary

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

var ErrWordNotFound = errors.New("word not found in the dictionary")

type Dictionary struct {
	entries    []Entry
	addChan    chan Entry
	removeChan chan string
}

func New() *Dictionary {
	d := &Dictionary{
		addChan:    make(chan Entry),
		removeChan: make(chan string),
	}
	d.loadFromFile()
	go d.listenChannels()
	return d
}

func (d *Dictionary) listenChannels() {
	for {
		select {
		case entry := <-d.addChan:
			d.handleAdd(entry)
		case word := <-d.removeChan:
			d.handleRemove(word)
		}
	}
}

func (d *Dictionary) Add(word, definition string) {
	entry := Entry{Word: word, Definition: definition}
	d.addChan <- entry
}

func (d *Dictionary) handleAdd(entry Entry) {
	d.entries = append(d.entries, entry)
	d.saveToFile()
}

func (d *Dictionary) Remove(word string) {
	d.removeChan <- word
}

func (d *Dictionary) handleRemove(word string) {
	for i, entry := range d.entries {
		if entry.Word == word {
			d.entries = append(d.entries[:i], d.entries[i+1:]...)
			d.saveToFile()
			return
		}
	}
}

func (d *Dictionary) Get(word string) (Entry, error) {
	for _, entry := range d.entries {
		if entry.Word == word {
			return entry, nil
		}
	}
	return Entry{}, ErrWordNotFound
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	sort.Slice(d.entries, func(i, j int) bool {
		return d.entries[i].Word < d.entries[j].Word
	})

	words := make([]string, len(d.entries))
	entriesMap := make(map[string]Entry)

	for i, entry := range d.entries {
		words[i] = entry.Word
		entriesMap[entry.Word] = entry
	}

	return words, entriesMap
}

func (d *Dictionary) saveToFile() {
	data, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("dictionary.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) loadFromFile() {
	data, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		panic(err)
	}
}
