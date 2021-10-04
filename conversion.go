package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 128
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //terjadi  integer overflow
	var (
		name           = "alvin"
		a       byte   = name[0]
		aString string = string(a)
	)

	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(aString)
}
