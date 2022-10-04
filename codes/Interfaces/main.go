package main

import "fmt"

type englishBot struct{}
type hindiBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eng := englishBot{}
	printGreeting(eng)
	hin := hindiBot{}
	printGreeting(hin)
}

func (e englishBot) getGreeting() string {
	return "hello there"
}

func (h hindiBot) getGreeting() string {
	return "Namaste"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
