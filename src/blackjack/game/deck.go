package game

import (
	"math/rand"
)

// Deck represents a single deck of playing cards, minus the Jokers.
type Deck []Card

var cards = []Card{
	Card{"A", "c"}, Card{"2", "c"}, Card{"3", "c"}, Card{"4", "c"},
	Card{"5", "c"}, Card{"6", "c"}, Card{"7", "c"}, Card{"8", "c"},
	Card{"9", "c"}, Card{"10", "c"}, Card{"J", "c"}, Card{"Q", "c"},
	Card{"K", "c"},
	Card{"A", "d"}, Card{"2", "d"}, Card{"3", "d"}, Card{"4", "d"},
	Card{"5", "d"}, Card{"6", "d"}, Card{"7", "d"}, Card{"8", "d"},
	Card{"9", "d"}, Card{"10", "d"}, Card{"J", "d"}, Card{"Q", "d"},
	Card{"K", "d"},
	Card{"A", "h"}, Card{"2", "h"}, Card{"3", "h"}, Card{"4", "h"},
	Card{"5", "h"}, Card{"6", "h"}, Card{"7", "h"}, Card{"8", "h"},
	Card{"9", "h"}, Card{"10", "h"}, Card{"J", "h"}, Card{"Q", "h"},
	Card{"K", "h"},
	Card{"A", "s"}, Card{"2", "s"}, Card{"3", "s"}, Card{"4", "s"},
	Card{"5", "s"}, Card{"6", "s"}, Card{"7", "s"}, Card{"8", "s"},
	Card{"9", "s"}, Card{"10", "s"}, Card{"J", "s"}, Card{"Q", "s"},
	Card{"K", "s"},
}

// NewDeck returns a new un-shuffled Deck.  Each Deck contains a private copy of the cards.
func NewDeck() Deck {
	deck := make([]Card, 52)
	copy(deck, cards)
	return deck
}

// Shuffle randomly reorders the cards inside the Deck
func (deck *Deck) Shuffle() {
	for i := range *deck {
		j := rand.Intn(i + 1)
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	}
}

// Deal "removes" the top card from the Deck and returns a pointer to that card
func (deck *Deck) Deal() *Card {
	card := (*deck)[len(*deck)-1]
	*deck = (*deck)[:len(*deck)-1]
	return &card
}

// Reset restores all cards to the Deck, but does not re/order them
func (deck *Deck) Reset() {
	*deck = (*deck)[:cap(*deck)]
}
