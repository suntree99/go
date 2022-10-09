package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {

	// Cara input property dengan (=)
	var budi Customer
	budi.Name = "Budi Darmawan"
	budi.Address = "Surabaya"
	budi.Age = 30

	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	// Cara key-value dengan (:)
	iwan := Customer {
		Name : "Iwan Setiawan",
		Address : "Bandung",
		Age : 35,
	}
	fmt.Println(iwan)

	// Cara langsung
	// tidak disarankan karena sangat tergantung urutan prototype
	wati := Customer { "Wati", "Tangerang", 25 }
	fmt.Println(wati)

}