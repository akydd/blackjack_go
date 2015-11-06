package main

import (
	"blackjack/game"
	"fmt"
	"log"
	"net/http"
)

var deck = game.NewDeck()
var playerHand, dealerHand game.Hand
var gameState = "NOT_RUNNING"

func start(w http.ResponseWriter, r *http.Request) {
	if gameState == "RUNNING" {
		fmt.Fprintf(w, "Game already in progress!\n\n")
	} else {
		gameState = "RUNNING"

		deck.Reset()
		deck.Shuffle()

		playerHand = game.NewHand()
		dealerHand = game.NewHand()

		playerHand.AddCard(deck.Deal())
		dealerShownCard := deck.Deal()
		dealerHand.AddCard(dealerShownCard)

		playerHand.AddCard(deck.Deal())
		dealerHand.AddCard(deck.Deal())

		fmt.Fprintf(w, "** INITIAL DEAL **\nYour Hand\n")

		printHand(w, playerHand)

		fmt.Fprintf(w, "Dealers shows: %q\n", dealerShownCard)
		fmt.Fprintf(w, "Dealers score is at least %v\n\n", dealerShownCard.Score())
	}
}

func printHand(w http.ResponseWriter, hand game.Hand) {
	fmt.Fprintf(w, "%q\n", hand)
	fmt.Fprintf(w, "Score: %v\n\n", hand.Score())
}

func hit(w http.ResponseWriter, r *http.Request) {
	if gameState == "NOT_RUNNING" {
		fmt.Fprintf(w, "Game not in progress!\n\n")
	} else {
		playerHand.AddCard(deck.Deal())
		score := playerHand.Score()

		fmt.Fprintf(w, "** HIT **\nYour Hand\n")
		printHand(w, playerHand)

		if score > 21 {
			fmt.Fprintf(w, "You lose!\n\n")
			gameState = "NOT_RUNNING"
		}
	}
}

func stand(w http.ResponseWriter, r *http.Request) {
	if gameState == "NOT_RUNNING" {
		fmt.Fprintf(w, "Game not in progress!\n\n")
	} else {
		fmt.Fprintf(w, "** STAND **\nDealer's Hand\n")
		printHand(w, dealerHand)

		dealerScore := dealerHand.Score()

		for dealerScore < 17 {
			dealerHand.AddCard(deck.Deal())
			dealerScore = dealerHand.Score()
			fmt.Fprintf(w, "** DEALER DRAW **\nDealer's Hand\n")
			printHand(w, dealerHand)
		}

		playerScore := playerHand.Score()
		fmt.Fprintf(w, "\n** FINAL SCORE **\nYou: %v\nDealer: %v\n", playerScore, dealerScore)
		if dealerScore > 21 {
			fmt.Fprintf(w, "You Win!")
		} else if dealerScore > playerScore {
			fmt.Fprintf(w, "You Lose!")
		} else {
			fmt.Fprintf(w, "You Win!")
		}
		gameState = "NOT_RUNNING"
	}
}

func main() {
	greeting()

	http.HandleFunc("/start", start)
	http.HandleFunc("/hit", hit)
	http.HandleFunc("/stand", stand)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greeting() {
	fmt.Println("Running on port 8080...")
	fmt.Println("")
	fmt.Println("** BLACKJACK! **")
	fmt.Println("")
	fmt.Println("Rules:")
	fmt.Println(" - Single player")
	fmt.Println(" - No betting")
	fmt.Println(" - Aces are always worth 11")
	fmt.Println("")
	fmt.Println("Instructions:")
	fmt.Println(" - '/start' starts a new game.  Once a game is started, it must be completed")
	fmt.Println(" - '/hit' to hit")
	fmt.Println(" - '/stand' to stand")
}
