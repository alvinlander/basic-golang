package main

import "fmt"

type Customers struct {
	Name, Address string
	Age           int
}

func (customer Customers) sayHello() {
	fmt.Println("Hello", customer.Name)
}
func main() {
	lans := Customers{
		Name: "Lander",
	}
	lans.sayHello()
}
