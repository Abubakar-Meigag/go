package main

import "fmt"

type contactInfo struct {
	email  string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	// The first way to create/declare a struct
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	// The second way to create/declare a struct
	var beko person
	fmt.Println(beko)
	fmt.Printf("%+v", beko)

	// How to update the properties of a struct
	var bob person
	bob.firstName = " Bob"
	bob.lastName = " Smith"
	fmt.Println(bob)
	fmt.Printf("%+v", bob)

	// Embedding a struct inside another struct
	jim := person{
		firstName: "Jim",
		lastName:  "Hill",
		contact: contactInfo{
			email:  "jim@gimail.com",
			zipCode: 12345,
		},
	}
	
	// Accessing the properties of a struct
	// jimPointer := &jim
	jim.updateName("James")
	jim.print()

	// Updating a property of an embedded struct
	jim.contact.updateEmail("james@gmail.com")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (c *contactInfo) updateEmail(newEmail string) {
	c.email = newEmail
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

