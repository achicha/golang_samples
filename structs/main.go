/* learn struct DS in golang. example will be the Person struct. */

package main

import "fmt"

// define contact, which will be embedded to the person's struct
type contactInfo struct {
	email   string
	zipCode int
}

// define person's struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// create a new struct of type person
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}

	peter := person{"Peter",
		"Olafson",
		contactInfo{email: "peter@gmail.com", zipCode: 95000}}

	var olaf person
	olaf.firstName = "Olaf"
	olaf.lastName = "Johanson"
	olaf.contact.email = "olaf@gmail.com"
	olaf.contact.zipCode = 96000

	// print output
	fmt.Printf("%+v \n", peter)

	// update Name using pointer
	alexPointer := &alex
	alexPointer.updateName("Alexey")
	alex.print()

	olaf.updateName("Olaff") // pointer shortcut, golang will handle with pointing inside.
	olaf.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
