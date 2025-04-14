package main

import (
	"os"
)

func (d deck) towrite(filename string) error{
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}