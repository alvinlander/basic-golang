package main

import (
	"fmt"
	"strconv"
)

//string conversion
func main() {
	bool, err := strconv.ParseBool("true")
	//bool, err := strconv.ParseBool("salah")
	if err == nil{
		fmt.Println(bool)
	}else{
		fmt.Println(err)
	}

	numb, err := strconv.ParseInt("100",10,32) // base : 2 binary, 8 octal, 10 decimal, 16 hexa decimal
	if err == nil{
		fmt.Println(numb)
	}else{
		fmt.Println(err)
	}

	fmt.Println("Integer to binary string",strconv.FormatInt(1000,2))

	fmt.Println("Integer to string : ", strconv.Itoa(30)) // atoi untuk sebaliknya
}