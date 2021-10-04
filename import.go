package main

import (
	"belajar-golang/helper"
	"fmt"
)
func main() {
	helper.SayHello("Lans")
	/**
	fmt.Println(helper.version)
	helper.sayGoodBye("Lans") Error karena huruf depan nama
	fungsi merupakan huruf kecil sehingga tidak dapat diakses
	oleh package lain
	*/
	fmt.Println(helper.Application)

}