package main

import "fmt"

func main() {
	cards := []string{"five of diamond", newCard()}
	cards2 := append(cards, "six of heart")
	fmt.Println(cards)
	fmt.Println(cards2)
}

func newCard() string {

	return "fifth of diamond"
}
