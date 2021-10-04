package helper

import "fmt"

var version int8= 1
var Application string= "Belajar Golang"
func SayHello(name string) {
	fmt.Println("Hello", name)
}
func sayGoodBye(name string){
	fmt.Println("Good bye", name)
}