package main

import "fmt"

func main() {

	name := "Budi"
	counter := 0

	increment := func() {

		name = "Iwan"
		fmt.Println("Increment")
		counter++

	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

}