package main

import "fmt"

func EndApp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Terjadi Error", msg)
	}
	fmt.Println("Aplikasi Selesai")
}
func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi Berjalan")
}
func main() {
	RunApp(false)
	// RunApp(true)
}
