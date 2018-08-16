package main

import (
	"fmt"
)

type contactInfo struct {
	address  string
	postcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Song",
		contact: contactInfo{
			address:  "123 Street",
			postcode: 3000,
		},
	}

	jim.updateName("Jimmy")

	jim.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
