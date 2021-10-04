package main

import (
	"flag"
	"fmt"
)

func main() {
	hostname := flag.String("host","localhost","put your hostname")
	username := flag.String("username","localhost","put your username")
	password := flag.String("password","localhost","put your password")

	flag.Parse()

	fmt.Println("hots =",*hostname)
	fmt.Println("username =",*username)
	fmt.Println("password =",*password)

}