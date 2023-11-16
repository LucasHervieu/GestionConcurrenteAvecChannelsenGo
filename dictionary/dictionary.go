// dictionary.go
package dictionary

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
)

// Déclaration du type Entry qui représente une entrée de dictionnaire
type Entry struct {
	Word       string `json:"word"`       // Mot
	Definition string `json:"definition"` // Définition
}

// Erreur personnalisée pour indiquer qu'un mot n'a pas été trouvé dans le dictionnaire
var ErrWordNotFound = errors.New("word not found in the dictionary")

// Déclaration du type Dictionary qui représente le dictionnaire
type Dictionary struct {
	entries []Entry // Liste des entrées du dictionnaire
}

// Fonction de création d'une nouvelle instance de Dictionary
func New() *Dictionary {
	d := &Dictionary{}
	d.loadFromFile() // Charge les données depuis le fichier
	return d
}

// Méthode Add ajoute un mot et sa définition au dictionnaire
func (d *Dictionary) Add(word, definition string) {
	entry := Entry{Word: word, Definition: definition}
	d.entries = append(d.entries, entry)
	d.saveToFile() // Enregistre les modifications dans le fichier
}

// Méthode Get récupère la définition d'un mot dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, error) {
	for _, entry := range d.entries {
		if entry.Word == word {
			return entry, nil
		}
	}
	return Entry{}, ErrWordNotFound
}

// Méthode Remove supprime un mot du dictionnaire
func (d *Dictionary) Remove(word string) {
	for i, entry := range d.entries {
		if entry.Word == word {
			d.entries = append(d.entries[:i], d.entries[i+1:]...)
			d.saveToFile() // Enregistre les modifications dans le fichier
			return
		}
	}
}

// Méthode List retourne la liste triée des mots et de leurs définitions
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

// Méthode saveToFile enregistre les entrées du dictionnaire dans un fichier JSON
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

// Méthode loadFromFile charge les données du fichier JSON dans le dictionnaire
func (d *Dictionary) loadFromFile() {
	data, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		// Si le fichier n'existe pas ou il y a une erreur de lecture, ignore simplement
		return
	}

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		panic(err)
	}
}
