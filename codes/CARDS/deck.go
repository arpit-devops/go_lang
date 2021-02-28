package main

import "fmt"

type deck []string

func newDeck() deck {
	var newCards deck
	cardVal1 := []string{"spade", "heart", "diamond"}
	cardVal2 := []string{"ace", "two", "three", "four"}

	for _, val1 := range cardVal1 {

		for _, val2 := range cardVal2 {

			newCards = append(newCards, val1+" of "+val2)

		}

	}
	return newCards
}

func deal(d deck, num int) (deck, deck) {

	return d[:num], d[num:]

}

func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)
	}

}
