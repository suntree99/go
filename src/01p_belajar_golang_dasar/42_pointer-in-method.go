package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() { // Tambahkan *
	man.Name = "Mr. " + man.Name
}

func main() {

	budi := Man{"Budi"}
	budi.Merried()
	fmt.Println(budi.Name)

}