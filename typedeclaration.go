package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool
	var alvin NoKTP = "337412xxxxx"
	var alvinStatus Married = false
	fmt.Println(alvin,alvinStatus)
}