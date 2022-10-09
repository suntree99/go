package main

import "regexp"
import "fmt"

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])a")
	fmt.Println(regex.MatchString("eka"))
	fmt.Println(regex.MatchString("eta"))
	fmt.Println(regex.MatchString("eDa"))

	search := regex.FindAllString("eka eko eda eta eya eki", -1)
	fmt.Println(search)
}