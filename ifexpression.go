package main

import "fmt"

func main() {
	name := "lans"

	if name == "lans" {
		// string concatenate
		fmt.Println(fmt.Sprintf("Halo %s", name))
	} else {
		fmt.Println("Bukan " + name)
	}
}
