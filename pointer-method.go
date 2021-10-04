package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	lans := Man{"Lander"}
	lans.Married()

	fmt.Println(lans.Name)
}
