package game

import (
	"reflect"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := NewDeck()

	if len(deck) != 52 {
		t.Error("expected 52 cards in the deck")
	}

	firstCard := deck[0]

	if firstCard.rank != "A" {
		t.Error("expected the first card rank to ba an Ace")
	}

	if firstCard.suite != "c" {
		t.Error("Expected the first card suite to be Clubs")
	}

	// Test that two decks are independent of each other
	deck2 := NewDeck()
	deck2[0], deck2[1] = deck2[1], deck2[0]
	if reflect.DeepEqual(deck[0], deck2[0]) {
		t.Error("Reordering elements of one deck should not affect another deck")
	}
}

func TestDeck_Deal(t *testing.T) {
	deck := NewDeck()
	topCard := deck[len(deck)-1]
	dealtCard := deck.Deal()
	newTopCard := deck[len(deck)-1]

	if reflect.DeepEqual(topCard, newTopCard) {
		t.Error("The card on the top of the Deck should have changed")
	}

	if !reflect.DeepEqual(topCard, *dealtCard) {
		t.Error("The top card should also be the dealt card")
	}

	if len(deck) != 51 {
		t.Error("The size of the remaining Deck should have decreased by 1")
	}
}

func TestDeck_Reset(t *testing.T) {
	deck := NewDeck()
	topCard := deck.Deal()
	deck.Deal()
	deck.Deal()

	deck.Reset()
	newTopCard := deck[51]

	if len(deck) != 52 {
		t.Error("Expected the deck to have all 52 cards again")
	}

	if !reflect.DeepEqual(*topCard, newTopCard) {
		t.Error("Expected the top card to be restored after resetting the deck")
	}
}
