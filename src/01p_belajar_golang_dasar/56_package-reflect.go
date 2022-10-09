package main

import "reflect"
import "fmt"

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string `required="true"`
	Description string
}

// fungsi untuk validasi
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {

	sample := Sample{"Budi"}

	var sampleType = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // kososng karena tidak ada tag min

	// contoh validasi
	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"Budi", ""} // true karena tidak dipasang tag required
	fmt.Println(IsValid(contoh))
	
}