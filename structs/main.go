package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	raja := person{
		firstName: "Raja",
		lastName:  "Sachin",
		contactInfo: contactInfo{
			email:   "sachinraja@gmail.com",
			pinCode: 560108,
		},
	}

	raja.updateFirstName("King")
	raja.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
