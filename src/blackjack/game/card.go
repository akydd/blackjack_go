package game

import (
	"fmt"
)

type Card struct {
	rank  string
	suite string
}

var scoreMap = map[string]int{
	"A":  11,
	"K":  10,
	"Q":  10,
	"J":  10,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
}

var rankNameMap = map[string]string{
	"A":  "Ace",
	"K":  "King",
	"Q":  "Queen",
	"J":  "Jack",
	"10": "10",
	"9":  "9",
	"8":  "8",
	"7":  "7",
	"6":  "6",
	"5":  "5",
	"4":  "4",
	"3":  "3",
	"2":  "2",
}

var suiteNameMap = map[string]string{
	"c": "Clubs",
	"h": "Hearts",
	"d": "Diamonds",
	"s": "Spades",
}

// Score returns the score of the card according to BlackJack rules.
// For simplicity, an Ace is always scored at 11.
func (c *Card) Score() int {
	return scoreMap[c.rank]
}

func (c *Card) String() string {
	return fmt.Sprintf("%v of %v", rankNameMap[c.rank], suiteNameMap[c.suite])
}
