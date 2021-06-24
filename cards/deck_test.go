package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Set up deck
	d := newDeck()

	// Calculate test values
	length := len(d)
	firstCard := d[0]
	lastCard := d[length-1]

	// Test attributes
	if length != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if firstCard != "Two of Spades" {
		t.Errorf("Expected first card to be 'Two of Spades', but got %v", firstCard)
	}

	if lastCard != "Jack of Clubs" {
		t.Errorf("Expected last card to be 'Jack of Clubs', but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewTestFromFile(t *testing.T) {
	// Setup
	os.Remove("_decktesting")
	d := deck{"Ace of Spades", "Queen of Hearts"}

	// Test saving to a file
	d.saveToFile("_decktesting")

	// Test loading from file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 2 {
		t.Errorf("Expected 2 cards in deck, got %v", len(loadedDeck))
	}

	// Cleanup
	os.Remove("_decktesting")
}
