package main

import "fmt"

func Logging() {
	fmt.Println("Selesai memanggil function")
}
/**

func RunApplication(value int) {
	fmt.Println("Run Application")
	res := 10 / value
	fmt.Println("Result", res)
	Logging()
}

*/
// Solution
func RunApplication(value int) {
	defer Logging()
	fmt.Println("Run Application")
	res := 10 / value
	fmt.Println("Result", res)

}
func main() {
	RunApplication(0)
}
