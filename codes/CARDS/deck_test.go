package main

import "testing"

func TestNewDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 129 {
		t.Errorf("Test fail in appropriate no of card, but got %v", len(cards))
	}

}
