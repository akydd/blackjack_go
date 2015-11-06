package game

// Hand represents a players hand of cards
type Hand []*Card

// NewHand creates an empty Hand and returns it
func NewHand() Hand {
	return make([]*Card, 0, 2)
}

// Score calculates and returns the total score for the Hand
func (hand *Hand) Score() int {
	score := 0

	for _, card := range *hand {
		score += card.Score()
	}

	return score
}

// AddCard adds a *Card to the Hand
func (hand *Hand) AddCard(card *Card) {
	*hand = append(*hand, card)
}
