package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToINdonesia(address *Address) { // Tambahkan *
	address.Country = "Indonesia"
}

func main() {

	var address10 = Address {
		City : "Subang",
		Province : "Jawa Barat",
		Country : "",
	}

	ChangeCountryToINdonesia(&address10) // Tambahkan &

	fmt.Println(address10)
	
}