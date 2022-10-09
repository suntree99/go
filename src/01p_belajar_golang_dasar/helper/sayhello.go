package helper

import "fmt"

var version = 1 // local scope karena diawali huruf kecil
var Application = "Belajar Golang Dasar" // global scope karena diawali huruf besar

func SayHello(name string) { // global scope karena diawali huruf besar
	fmt.Println("Hello", name)
}

func sayGoodBye (name string) { // local scope karena diawali huruf kecil
	fmt.Println("Good Bye", name)
}