package game

import "testing"

func TestHand(t *testing.T) {
	hand := NewHand()

	if len(hand) != 0 {
		t.Error("expected an initial length of 0")
	}

	if cap(hand) != 2 {
		t.Error("expected an initial capacity of 2")
	}
}

func TestHand_addCard(t *testing.T) {
	hand := NewHand();
	card1 := Card{"A", "c"}
	card2 := Card{"3", "s"}

	(&hand).AddCard(&card1);
	firstCard := hand[0]

	if len(hand) != 1 {
		t.Error("expected the length of the Hand to have increased to 1")
	}

	if firstCard.rank != "A" {
		t.Error("expected the first card rank to ba an Ace")
	}

	if firstCard.suite != "c" {
		t.Error("Expected the first card suite to be Clubs")
	}

	(&hand).AddCard(&card2)
	secondCard := hand[1]

	if len(hand) != 1 {
		t.Error("expected the length of the Hand to have increased to 1")
	}

	if secondCard.rank != "3" {
		t.Error("expected the second card rank to ba a 3")
	}

	if secondCard.suite != "s" {
		t.Error("Expected the second card suite to be Spades")
	}
}

func TestHand_Score(t *testing.T) {
	hand := NewHand()

	if (&hand).Score() != 0 {
		t.Error("expected an empty hand to have a score of 0")
	}

	card1 := Card{"A", "c"}
	card2 := Card{"3", "s"}

	(&hand).AddCard(&card1)
	(&hand).AddCard(&card2)

	if (&hand).Score() != 14 {
		t.Error("expected the hand [Ac, 3s] to have a score of 14")
	}
}