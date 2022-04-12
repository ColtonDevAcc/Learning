package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
		os.Exit(1)
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
		os.Exit(1)
	}
}
