package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of \"%s\", but got %s", "Ace of Spades", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of \"%s\", but got %s", "Four of Clubs", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck is 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()

	d1, d2 := deal(d, 5)

	if len(d1) != 5 {
		t.Errorf("Expected length of 5 in d1, but got %v", len(d1))
	}

	if len(d2) != 11 {
		t.Errorf("Expected length of 11 in d1, but got %v", len(d2))
	}
}
