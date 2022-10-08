package main

import "fmt"

func getGoodBy(name string) string {
	return "Good By " + name
}

func main() {

	sayGoodBy := getGoodBy
	result := sayGoodBy("Budi")
	fmt.Println(result)

}