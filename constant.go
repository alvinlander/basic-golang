package main

import "fmt"

func main() {
	const firstName string = "Alvin"
	const lastname = "Lander"
	fmt.Println(firstName)
	fmt.Println(lastname)

	// type 2
	const (
		a = 1
		b = 2
	)
	fmt.Println(a + b)
	// note : value tidak dapat diubah

}
