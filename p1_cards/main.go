package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()

	fmt.Println(card)

	cards := []string{"Ace of Spades", newCard(), newCard()}

	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
