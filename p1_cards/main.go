package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()

	// fmt.Println(card)

	// cards := []string{"Ace of Spades", newCard()}
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)
	// fmt.Println(cards[1])
	// fmt.Println(cards[0:2])
	// fmt.Println(cards[:2])

	// for i := 0; i < len(cards); i++ {
	// 	fmt.Println(cards[i])
	// }

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()
}

func newCard() string {
	return "Two of Spades"
}
