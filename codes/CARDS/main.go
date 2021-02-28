package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {

	return "fifth of diamond"
}
