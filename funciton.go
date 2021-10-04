package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloWithParam(firstName string, lastName string) {
	fmt.Println(fmt.Sprintf("Hello %s %s", firstName, lastName))
}

// with return
func sayHelloR() string {
	return "Hello"
}

func sayHelloWithParamR(firstName string, lastName string) string {
	return fmt.Sprintf("Hello %s %s", firstName, lastName)
}

// multiple return
func getFullname() (string, string) {
	return "Alvin", "Lander"
}

// return named value
func getName() (firstName, middleName, lastName string) {
	firstName = "First String"
	middleName = "Middle String"
	lastName = "Last String"
	return
}
func main() {
	sayHello()
	sayHelloWithParam("Alvin", "Lander")
	// with return
	fmt.Println(sayHelloR())
	fmt.Println(sayHelloWithParamR("Lander", "Alvin"))
	// multiple return values
	firstName, lastname := getFullname()
	// if need just fiirstname or lastname
	fmt.Println(firstName, lastname)
	fName, _ := getFullname()
	fmt.Println(fName)
	// return named values
	_, _, thirdName := getName()
	fmt.Println(thirdName)
}
