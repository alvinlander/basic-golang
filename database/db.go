package database

import "fmt"

var connection string

//fungsi untuk meng inisasi pertama kali tanpa harus di jalankan
func init(){
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string{
	return connection
}