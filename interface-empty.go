package main

import "fmt"

/**
Kegunaan interface kosong agar
tidak mementingkan tipe data.
Jadi bisa menggunakan semua tipe data
*/
func Ups(value int) interface{} {
	if value == 1 {
		return value
	} else if value == 2 {
		return true
	} else {
		return "ok"
	}
}
func main() {
	fmt.Println(Ups(1))
	fmt.Println(Ups(2))
	fmt.Println(Ups(3))
}
