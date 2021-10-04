package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains :",strings.Contains("Alvin Lander","Land"))
	fmt.Println("Split (slice):",strings.Split("Alvin Lander"," "))
	fmt.Println("Lower :",strings.ToLower("HaYo"))
	fmt.Println("Upper :",strings.ToUpper("alvin"))
	fmt.Println("Title :",strings.ToTitle("lander"))
	fmt.Println("Trim :",strings.Trim("     Lans      "," "))  //memotong spasi yang di kiri kanan string
	fmt.Println("Replace all :", strings.ReplaceAll("Alvin Alvin Lander","Alvin","Lans"))

}