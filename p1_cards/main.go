package main

func main() {
	cards := newDeckFromFile("myCards.txt")
	cards.print()
}
