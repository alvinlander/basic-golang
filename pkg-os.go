package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)
	fmt.Println()
	if len(args) >= 1 {
		for i,_:= range args{
			fmt.Println(args[i])
		}
	}
	name, err := os.Hostname()
	if(err == nil){
		fmt.Println("Hostname",name)
	}else{
		fmt.Println("Error")
	}
}