package main

import (
	"os"
	"testing"
)

// run package means run all tests that are defined to this package
// run file tests means run all tests in this file

// must have parameter of t *testing.T
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// this is a formatted string
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %d", len(d))
	}

	os.Remove("_decktesting")
}