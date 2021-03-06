package main

import "fmt"

func main() {

	colours := map[string]string{

		"RED":   "R",
		"BLUE":  "B",
		"GREEN": "G",
	}

	printMap(colours)
}

func printMap(m map[string]string) {

	for key, value := range m {

		fmt.Println(key, " : ", value)
	}
}
