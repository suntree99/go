package main

import "fmt"
import "strings"

func main() {

	fmt.Println(strings.Contains("Budi Darmawan", "Budi"))
	fmt.Println(strings.Contains("Budi Darmawan", "Iwan"))

	fmt.Println(strings.Split("Budi Darmawan", " "))

	fmt.Println(strings.ToLower("Budi Darmawan"))
	fmt.Println(strings.ToUpper("Budi Darmawan"))
	fmt.Println(strings.ToTitle("Budi Darmawan"))

	fmt.Println(strings.Trim("     Budi Darmawan     ", " "))

	fmt.Println(strings.ReplaceAll("Budi Budi Budi Iwan Wati", "Budi", "Darmawan"))

}