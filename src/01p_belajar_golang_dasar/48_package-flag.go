package main

import "flag"
import "fmt"

func main() {

	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse() // untuk memparsing data masukan

	fmt.Println("Host : ", *host)
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

	// go run .\48_package-flag.go -username=budi -number=budi
	// Akan menghasilkan error dan instruksi pengisian

}