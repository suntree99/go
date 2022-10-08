package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// Pass by value
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1
	address2.City = "Subang" // perubahan address2 TIDAK MENGUBAH reference (address1)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("")

	// Pass by reference -> Pointer (&)
	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3
	address4.City = "Lamongan" // perubahan address4 MENGUBAH reference (address3)

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("")

	// Assign nilai baru TIDAK MENGUBAH reference
	address5 := Address{"Tangerang", "Banten", "Indonesia"}
	address6 := &address5
	address6 = &Address{"Yogyakarta", "DIY", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println("")

	// Assign nilai baru MENGUBAH reference dengan (*)
	address7 := Address{"Pontianak", "Kalimantan Barat", "Indonesia"}
	address8 := &address7
	*address8 = Address{"Samarinda", "Kalimantan Timur", "Indonesia"}

	fmt.Println(address7)
	fmt.Println(address8)
	fmt.Println("")

	var address9 *Address = new(Address)
	address9.City = "Jakarta"
	fmt.Println(address9)
	
}