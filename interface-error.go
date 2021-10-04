package main

import (
	"errors"
	"fmt"
)

func Pembagi(value, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi dengan NOL")
	} else {
		return value / pembagi, nil
	}
}
func main() {
	error := errors.New("ups error")
	fmt.Println(error)
	res, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("Hasil", res)
	} else {
		fmt.Println("ERROR", err.Error())
	}
}
