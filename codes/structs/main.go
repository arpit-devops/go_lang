package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{
		firstName: "alex",
		lastName:  "ficher",
		contactInfo: contactInfo{
			email:   "alex@ficher.com",
			pincode: 400040,
		},
	}
	alexPointer := &alex
	alexPointer.editName("alexa")
	alex.print()
}

func (pointerToAlex *person) editName(newFirstName string) {

	(*pointerToAlex).firstName = newFirstName
}

func (p person) print() {

	fmt.Printf("%+v", p)
}
