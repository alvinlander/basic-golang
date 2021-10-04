package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 100
	var c = a * b
	fmt.Println(c)

	// augmented assignments
	d := 0
	d += 10
	d -= 5
	d *= 2
	d /= 2
	d %= 2
	fmt.Println(d)

	//unary operator
	// a++ -> a = a+1
	// a-- -> a = a-1
	// -a -> -a (negatif a)
	// a -> a(positif a)
	// !a -> kebalikan boolean

}
