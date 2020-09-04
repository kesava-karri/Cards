package main

import "testing"

func TestNewDeck(t *testing.T) { //t is the test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}
}
