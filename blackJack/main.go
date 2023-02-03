package main

import (
	"fmt"

	"github.com/Eric-lab-star/deck"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]

		fmt.Println(card)

	}
}
