package game

import(
    "testing"
    "reflect"
)

func TestDeck(t *testing.T) {
    deck := NewDeck()
    
    if len(deck) != 52 {
        t.Error("expected 52 cards in the deck")
    }
    
    firstCard := deck[0];
    
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
    deck := NewDeck();
    lastCard := deck[len(deck) - 1]
    dealtCard := (&deck).Deal()
    newLastCard := deck[len(deck) - 1]
    
    if reflect.DeepEqual(lastCard, newLastCard) {
        t.Error("The card on the top of the Deck should have changed")
    }
    
    if !reflect.DeepEqual(lastCard, *dealtCard) {
        t.Error("The top card should also be the dealt card")
    }
}