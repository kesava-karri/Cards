package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //t is the test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected Four of Hearts, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of 16, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
