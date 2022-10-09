package main

import "os"
import "fmt"

func main() {

	// Args
	args := os.Args
	fmt.Println("Arguments : ")
	fmt.Println(args)

	// Hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error ; ", err.Error())
	}

	// Environment variables
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	// export APP_USERNAME=root
	// export APP_PASSWORD=root

}