package game

import "testing"

func TestCard_Score(t *testing.T) {
	card := Card{"A", "c"}
	if card.Score() != 11 {
		t.Error("expected 11 for an Ace")
	}

	card.rank = "K"
	if card.Score() != 10 {
		t.Error("expected 10 for a King")
	}

	card.rank = "Q"
	if card.Score() != 10 {
		t.Error("expected 10 for a Queen")
	}

	card.rank = "J"
	if card.Score() != 10 {
		t.Error("expected 10 for a Jack")
	}
}
