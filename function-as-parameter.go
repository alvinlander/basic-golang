package main

import "fmt"

//  function as a parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	// func as params
	sayHelloWithFilter("Anjing", spamFilter)
}
