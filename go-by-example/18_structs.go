package main

import "fmt"


// declares a person struct (class)
type person struct {
	name string
	age  int
}

// constructor returns pointer to struct of type person
func NewPerson(name string, age int) *person {
	p := person{
		name: name,
		age: age,
	}

	return &p
}

func main() {
	fmt.Println(person{"Ibtida", 20})	// implicit

	fmt.Println(person{			// named fields 
		name: "Bob",
		age: 20,
	})

	fmt.Println(NewPerson("Joseph", 32))	// using constructor

	instance := NewPerson("Tomathy", 15)	// instance is a *person pointer

	fmt.Println("name:", instance.name) // Go automatically dereferences w/ dot operator
	fmt.Println("age:", instance.age)
}
