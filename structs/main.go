package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "alex.anderson@email.com",
			zip:   3333,
		},
	}
	alex.print()
	var mesut person
	mesut.firstName = "Mesut"
	mesut.lastName = "Yilmaz"
	mesut.contactInfo.email = "mesutyilmaz@email.com"
	mesut.contactInfo.zip = 33332
	mesutPointer := &mesut
	mesutPointer.updateName("mesut")
	mesut.updateLastName("yilmazzz")
	mesut.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)

}
