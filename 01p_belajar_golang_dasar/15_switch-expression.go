package main

import "fmt"

func main() {

	name := "Darmawan"

	switch name {
	case "Budi" :
		fmt.Println("Hello Budi")
	case "Darmawan" :
		fmt.Println("Hello Darmawan")
	default :
		fmt.Println("Hi, boleh kenalan?")
	}

	// short expression
	switch length := len(name) ; length > 5 {
	case true :
		fmt.Println("Nama terlalu panjang")
	case false :
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa ekspresi (seperti if)
	length := len(name)
	switch {
	case length > 10 :
		fmt.Println("Nama terlalu panjang")
	case length > 5 :
		fmt.Println("Nama lumayan panjang")
	default :
		fmt.Println("Nama sudah benar")
	}
}