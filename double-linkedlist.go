package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Alvin") // Alvin
	data.PushBack("Lander") // Alvin -> Lander
	data.PushFront("Lans") // Lans -> Alvin -> Lander

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println()
	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
	// reverse
	fmt.Println()
	fmt.Println("Reverse")
	for e := data.Back(); e != nil; e = e.Prev(){
		fmt.Println(e.Value)
	}
}
