package main

import "fmt"

type deck []string


func (d deck) print() {
	for index,card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	var cards deck
	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	
	for _, suit := range cardSuit{
		for _, value := range cardValue{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
} 