package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorilaRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorilaRecursive(value - 1)
	}
}

func main() {

	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5*4*3*2*1)

	recursive := factorilaRecursive(5)
	fmt.Println(recursive)

}