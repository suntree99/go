package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // tetap akan dieksekusi
	fmt.Println("Run Appliaction") // dieksekusi
	result := 10/value // error
	fmt.Println("Result", result) // tidak dieksekusi
}

func main() {

	runApplication(0)

}