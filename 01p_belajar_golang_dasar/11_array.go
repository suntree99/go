package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Budi"
	names[1] = "Darmawan"
	names[2] = "Suntree"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var value = [3]int {
		90,
		95,
		80,
	}

	fmt.Println(value)

	var lagi [10]string
	fmt.Println(len(names))
	fmt.Println(len(value))
	fmt.Println(len(lagi))

}
