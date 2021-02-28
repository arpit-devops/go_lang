package main

import "math/rand"

func (d deck) shuffle() deck {

	for i := range d {
		newPos := rand.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]

	}
	return d
}
