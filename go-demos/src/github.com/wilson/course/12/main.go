package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//	Methods
//
// Greeting
// func (p Person) Greeting() (string, int) {
// 	return "Hello, my name is " + p.firstName + " and I'm", p.age

// }

// Using strconv: convert integer at string value
func (p Person) Greeting() string {
	return "Hello, my name is " + p.firstName + " and I'm " + strconv.Itoa(p.age)

}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(newLastName string) {

	switch p.gender {
	case "F":
		p.lastName = newLastName
	default:
		break
	}

}

func main() {

	// Init person using struct
	person1 := Person{firstName: "Wilson",
		lastName: "Costa",
		city:     "São Paulo",
		gender:   "M",
		age:      35}

	//  Alternative
	person2 := Person{"Jessica", "Santo", "São Paulo", "F", 31}

	// Show
	fmt.Println(person1, "\n", person2)

	fmt.Println(person2.Greeting())

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.age)

	person2.getMarried("Costa")
	fmt.Println(person2)

}
