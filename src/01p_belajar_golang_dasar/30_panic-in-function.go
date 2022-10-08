package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp() // tetap akan dieksekusi
	if error {
		panic("APLIKASI ERROR") // panic message dieksekusi saat error
	}
	fmt.Println("Aplikasi berjalan") // tidak akan dieksekusi
}

func main() {

	runApp(true)
	fmt.Println("Program berjalan") // tidak akan dieksekusi saat terjadi error

}