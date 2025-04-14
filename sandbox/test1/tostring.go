package main

import (
	"strings"
)

func (d deck) toString() string {
	deckStringSlice := []string(d)
	joinedString :=  strings.Join(deckStringSlice, ", ")
	return joinedString
}

