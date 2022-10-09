package main

import "container/list"
import "fmt"

func main() {

	data := list.New()

	data.PushBack("Budi")
	data.PushBack("Darmawan")
	data.PushFront("Suntree")

	// menampilkan data pertama atau terakhir
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// menampilkan data ujung (nil)
	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}