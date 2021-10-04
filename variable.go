package main

import "fmt"

func main(){
	// type 1
	var a string //value yang di input harus sesuai dengan deklarasi variabel
	a = "Alvin Lander"
	fmt.Println(a)
	// type 2
	var b = "alvin lander" //tipe data terdeteksi sendiri tanpa harus memasukan tipe data
	fmt.Println(b)
	// type 3
	c := "lander alvin"
	c = "lander" //untuk mengubah value tidak usah pake :
	fmt.Println(c)
	// type 4
	var(
		firstName = "Alvin"
		lastName = "Lander"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}