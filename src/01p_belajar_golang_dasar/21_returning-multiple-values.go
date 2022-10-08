package main

import "fmt"

func getFullName() (string, string, int) {
	return "Budi", "Darmawan", 33
}

func main() {

	firstName, lastName, age := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	namaDepan, namaBelakang, _ := getFullName()
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}