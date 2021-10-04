package main

import "fmt"

func main(){

	// Tipe data integer:
	// int8 = -128 - 127
	// int16 = -32768 - 32767
	// int32 = -2147483658 - 2147483657
	// int64 = -92233721036854775808 - 92233721036854775807
	// uint8 = 0 - 255
	// uint16 = 0 - 65535
	// uint32 = 0 - 4294967295
	// uint64 = 0 - 18446744073709551615
	//tipe data float:
	// float32 = 1.18x10^-38 - 3.4x10^38
	// float64 = 2.23x10^-308 - 1.18x10^308
	// complex64 = complex number with float32
	// complex128 = complex number with float64
	// alias
	// byte = uint8
	// rune = int32
	// int = min32
	// uint = minuint32
	// *note : secara default ketika kita masukan
	// int/float saja akan mengassign sesuai
	// sistem operasi kita misal win10 64bit(maka int64)
	fmt.Println("Satu",1)
	fmt.Println("Dua",2)
}