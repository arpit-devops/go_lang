package main

import "fmt"

type bot1 interface {
	getChat1() string
}

type englishBot1 struct{}

type hindiBot1 struct{}

func (e englishBot1) getChat1() string {

	return "hiii there"
}

func (e hindiBot1) getChat1() string {

	return "Namaste"
}

func printChat1(b bot1) {

	fmt.Println(b.getChat1())
}

func main() {

	eng := englishBot1{}
	hindi := hindiBot1{}
	printChat1(eng)
	printChat1(hindi)

}
