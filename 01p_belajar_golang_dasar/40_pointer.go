package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// Pass by value
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1
	address2.City = "Subang"

	fmt.Println(address1)
	fmt.Println(address2)

	// Pass by reference (Pointer)
	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3
	address4.City = "Lamongan"

	fmt.Println(address3)
	fmt.Println(address4)

	// Assign nilai baru
	// address4 = Address{"Bandung", "Jawa Barat", "Indonesia"}

	// fmt.Println(address3)
	// fmt.Println(address4)
}