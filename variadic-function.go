package main

import "fmt"

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	// variadic function
	sum := sumAll(10, 12, 13)
	fmt.Println(sum)
	// if has a sluce using variadic function
	numbers := []int{10, 10, 10}
	tot := sumAll(numbers...)
	fmt.Println(tot)
}
