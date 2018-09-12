package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Excepted deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaces, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"

	// remove previous testing files (if any)
	os.Remove(fileName)

	// save to file
	deck := newDeck()
	deck.saveToFile(fileName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("File is not exists on harddrive, after saving. path: %v", fileName)
	}

	// load file from disk
	loadedDeck := newDeckFromFile(fileName)

	if deck.toString() != loadedDeck.toString() {
		t.Errorf("Read from file failed.")
	}

	// remove previous testing files (if any)
	os.Remove(fileName)

}