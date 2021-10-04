package main

import "fmt"

type Address struct {
	City, Province, Country string
}
//pointer function
func changeProvince(address *Address){
	address.Province = "Jawa Timur"
}
func main() {
	address1 := Address{
		City:     "Cirebon",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	/**
	Jika melakukan seperti ini maka akan
	mencopy value dari address1 sehingga,
	jika ada perubahaan pada address2 maka
	address1 tidak ikut berubah
	{Cirebon Jawa Barat Indonesia}
	{Bandung Jawa Barat Indonesia}
	*/
	//address2 := address1

	/**
	Pointer is Solution !
	maka harus melakukan pass by references.
	yaitu arahkan address2 ke pointer yang sama dengan address1
	{Bandung Jawa Barat Indonesia}
	&{Bandung Jawa Barat Indonesia}
	*/
	address2 := &address1
	address2.City = "Bandung"

	/**
	Ada 1 teknik pointer lagi jika addres2 membuat data baru.
	problemnya ketika address2 membuat data baru, address1 tidak ikut berubah
	{Bandung Jawa Barat Indonesia}
	&{Jakarta DKI Jakarta Indonesia}
	*/
	//address2 = &Address{"Jakarta","DKI Jakarta","Indonesia"}

	/**
	Solution!
	Maka solusinya agar semua variabel yg mengarah ke address1 ikut berubah.
	Maka harus menggunakan operator * di depan variabel
	{Jakarta DKI Jakarta Indonesia}
	&{Jakarta DKI Jakarta Indonesia}
	&{Jakarta DKI Jakarta Indonesia}

	*/
	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	//membuat address3 untuk validasi
	address3 := &address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	//Membuat pointer dengan data kosong
	address4 := new(Address)
	fmt.Println(address4)
	address4.City = "Malang"
	fmt.Println(address4)

	//Pointer function
	address5 := Address{"Malang", "Jawa Barat", "Indonesia"}
	fmt.Println(address5)
	changeProvince(&address5)
	fmt.Println(address5)
}
