package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	var budi Customer
	budi.Name = "Budi Darmawan"
	budi.Address = "Surabaya"
	budi.Age = 33

	budi.sayHi("Iwan")

}