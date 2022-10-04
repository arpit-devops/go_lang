package main

import "fmt"

func main() {
	cards := []string{"five of diamond", newCard()}
	cards = append(cards, "six of heart")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {

	return "fifth of diamond"
}
