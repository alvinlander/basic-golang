package main

import "fmt"

func main() {
	name1 := "Eko"
	name2 := "Budi"

	compare := name1 == name2
	fmt.Println(compare)

	value1 := 5
	value2 := 3
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
