package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.GetName())
}

// Person struct
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Animal struct
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var budi Person
	budi.Name = "Budi Darmawan"

	SayHello(budi)

	cat := Animal {
		Name : "Push",
	}

	SayHello(cat)

}