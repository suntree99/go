package main

import "container/ring"
import "strconv"
import "fmt"

func main() {

	// var data *ring.Ring = ring.New(5)
	data := ring.New(5)

	// cara input manual
	// data.Value = "Budi"
	// var data2 = data.Next();
	// data2.Value = "Darmawan"

	// cara input dengan iterasi
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// mencetak seluruh data ring
	data.Do(func(value interface{}){
		fmt.Println(value)
	})
}