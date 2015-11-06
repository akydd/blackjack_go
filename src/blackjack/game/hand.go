package game

type Hand []*Card

// Create an empty Hand
func NewHand() Hand {
	return make([]*Card, 0, 2)
}

// Return the total score for the Hand
func (hand *Hand) Score() int {
	score := 0

	for _, card := range *hand {
		score += card.Score()
	}

	return score
}

// Add a *Card to the Hand
func (hand *Hand) AddCard(card *Card) {
	*hand = append(*hand, card)
}
