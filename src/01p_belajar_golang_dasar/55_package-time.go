package main

import "time"
import "fmt"

func main() {

	// waktu saat ini
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// mengeset tanggal
	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	// parsing waktu
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)

}