package main

import "fmt"

func main() {
	// Buat array manual
	var names [3]string
	names[0] = "Alvin"
	names[1] = "Lander"
	names[2] = "Lans"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := []int{
		90, 80, 88,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))
}
