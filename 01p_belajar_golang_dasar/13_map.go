package main

import "fmt"

func main() {

	person := map[string]string {
		"name" : "Budi",
		"address" : "Surabaya",
	}

	person["title"] = "Programmmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["titile"] = "Belajar Go-Lang"
	book["author"] = "Budi"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}