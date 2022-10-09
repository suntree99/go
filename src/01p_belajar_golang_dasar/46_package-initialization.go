package main

import "01p_belajar_golang_dasar/database"
// import _ "01p_belajar_golang_dasar/database" // blank identifier (_)
import "fmt"

func main() {

	result := database.GetDatabase()
	fmt.Println(result)

}