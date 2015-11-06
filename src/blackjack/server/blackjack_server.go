package main

import(
    "fmt"
    "log"
    "net/http"
    "blackjack/game"
)

var deck = game.NewDeck()
var playerHand, dealerHand game.Hand
var gameState = "NOT_RUNNING"


func start(w http.ResponseWriter, r *http.Request) {
    if (gameState == "RUNNING") {
        fmt.Fprintf(w, "Game already in progress!\n\n")
    } else {
        gameState = "RUNNING"

        (&deck).Reset()
        (&deck).Shuffle()

        playerHand = game.NewHand()
        dealerHand = game.NewHand()

        (&playerHand).AddCard((&deck).Deal())
        dealerShownCard := (&deck).Deal()
        (&dealerHand).AddCard(dealerShownCard)

        (&playerHand).AddCard((&deck).Deal())
        (&dealerHand).AddCard((&deck).Deal())

        fmt.Fprintf(w, "** INITIAL DEAL **\n")
        fmt.Fprintf(w, "Your hand: %q\n", playerHand)
        fmt.Fprintf(w, "Your hand score is %v\n\n", (&playerHand).Score())
        fmt.Fprintf(w, "Dealers shows: %q\n", dealerShownCard)
        fmt.Fprintf(w, "Dealers score is at least %v\n\n", dealerShownCard.Score())
    }
}

func hit(w http.ResponseWriter, r *http.Request) {
    if (gameState == "NOT_RUNNING") {
        fmt.Fprintf(w, "Game not in progress!\n\n")
    } else {
        (&playerHand).AddCard((&deck).Deal())
        score := (&playerHand).Score()

        fmt.Fprintf(w, "** HIT **\n")
        fmt.Fprintf(w, "Your hand: %q\n", playerHand)
        fmt.Fprintf(w, "Your hand score is %v\n\n", score)

        if (score > 21) {
            fmt.Fprintf(w, "You lose!\n\n")
            gameState = "NOT_RUNNING"
        }
    }
}

func stand(w http.ResponseWriter, r *http.Request) {
    if (gameState == "NOT_RUNNING") {
        fmt.Fprintf(w, "Game not in progress!\n\n")
    } else {
        fmt.Fprintf(w, "** STAND **\n")

        dealerScore := (&dealerHand).Score()
        fmt.Fprintf(w, "The dealer reveals hand %q\n", dealerHand)
        fmt.Fprintf(w, "The dealers score is %v\n", dealerScore)

        // TODO - play to the end
        for ; dealerScore < 17; {
            dealerCard := (&deck).Deal()
            (&dealerHand).AddCard(dealerCard)
            dealerScore = (&dealerHand).Score()

            fmt.Fprintf(w, "Dealer draws %q, bring the score to %v\n", dealerCard, dealerScore)
        }
        gameState = "NOT_RUNNING"
    }
}

func main() {
    http.HandleFunc("/start", start)
    http.HandleFunc("/hit", hit)
    http.HandleFunc("/stand", stand)

    log.Fatal(http.ListenAndServe(":8080", nil))
}