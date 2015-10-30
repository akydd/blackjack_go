package game

type Hand []Card

// Return the total score for the Hand
func (hand *Hand) Score() int {
    score := 0

    for _, card := range *hand {
        score += card.Score()
    }

    return score
}

// Create an empty Hand
func NewHand() Hand {
    return make([]Card, 2)
}

//func (hand *Hand) addToHand(card *Card) {
//    append(*hand, *card)
//}