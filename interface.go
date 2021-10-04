package main

import "fmt"

/**
Interface untuk membuat kontrak
*/
type HasName interface {
	getName() string
}

func Hello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}
func main() {
	var lans Person
	lans.Name = "Lander"
	Hello(lans)
	cat := Animal{
		Name: "Copa",
	}
	Hello(cat)
}
