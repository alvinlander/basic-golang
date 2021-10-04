package main

import (
	"fmt"
	"strconv"
)

func main() {

	// for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	// for range

	names := []string{"Alvin", "Lander"}

	for index, value := range names {
		fmt.Println(fmt.Sprintf("Index %s value %s ", strconv.Itoa(index), value))
	}

	for _, value := range names {
		fmt.Println(fmt.Sprintf("Values %s", value))
	}

	person := make(map[string]interface{})

	person["name"] = "Alvin"
	person["Age"] = 22

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}
}
