package main

import "strconv"
import "fmt"

func main() {

	// String to other type
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
	
	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)

	// Other Type to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueString := strconv.Itoa(45)
	fmt.Println(valueString)

}