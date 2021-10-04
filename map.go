package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alvin",
		"address": "Cirebon",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Software Engineer"
	fmt.Println(person)
	fmt.Println(person["title"])

	// make empty map

	book := make(map[string]interface{})
	book["title"] = "Buku Golang"
	book["pages"] = 256
	fmt.Println(book)
	book["ups"] = "Salah"

	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(len(book))

	// Membuat map di dalam map

	mahasiswa := map[string]map[string]string{
		"1001":map[string]string{
			"Nama": "Alvin",
			"Address": "Cirebon",
		},
		"1002":map[string]string{
			"Nama": "Lander",
			"Address": "Bandung",
		},
	}
	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1001"]["Nama"])
}
