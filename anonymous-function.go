package main

import "fmt"

// type declaration func
type BlackList func(string) bool

func userRegister(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	blackList := func(name string) bool {
		return name == "Anjing"
	}
	userRegister("Alvin", blackList)
}
