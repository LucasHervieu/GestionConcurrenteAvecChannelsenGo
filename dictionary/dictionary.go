// dictionary.go
package dictionary

import "errors"

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, errors.New("Word not found in the dictionary")
	}
	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() map[string]Entry {
	return d.entries
}
