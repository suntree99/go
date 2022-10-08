package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var noKtpBudi NoKTP = "3603090710890001"
	var marriedStatus Married = true
	fmt.Println(noKtpBudi)
	fmt.Println(marriedStatus)

}