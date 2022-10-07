package main

import "fmt"

func main() {

	var name = "Darmawan"

	if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Darmawan" {
		fmt.Println("Hello Darmawan")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	// short expression
	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}