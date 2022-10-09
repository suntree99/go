package main

import "01p_belajar_golang_dasar/helper"
import "fmt"

func main() {

	helper.SayHello("Budi")
	// helper.sayGoodBye("Budi") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

}