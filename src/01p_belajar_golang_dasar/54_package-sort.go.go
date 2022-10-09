package main

import "sort"
import "fmt"

type User struct {
	Name string
	Age int
}

// membuat struct menjadi array
type UserSlice []User

// membuat interface untuk Len()
func (value UserSlice) Len() int {
	return len(value)
}

// membuat interface untuk Less(i, j int)
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age <  value[j].Age
}

// membuat interface untuk Swap(i, j int)
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []User {
		{"Budi", 30},
		{"Iwan", 35},
		{"Wati", 25},
		{"Ani", 20},
		{"Andi", 15},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}