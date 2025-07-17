package main

import "fmt"

type person struct {
	firstName string
	lastName  string
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

}