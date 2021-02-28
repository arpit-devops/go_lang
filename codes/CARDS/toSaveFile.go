package main

import (
	"io/ioutil"
	"strings"
)

func (d deck) toString() string {

	deckStringSlice := []string(d)
	singleString := strings.Join(deckStringSlice, ",")
	return singleString
}

func (d deck) toWriteFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}
