package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {

	// creating map
	// m := make(map[string]string)

	// setting an element
	// m["fname"] = "Faizan"
	// m["lname"] = "Shaikh"

	// get an element
	// IMP: if key does not exists in the map then it returns zero value
	// fmt.Println(m["fname"])
	// fmt.Println(m["lname"])

	// delete(m, "fname")
	// clear(m)

	// fmt.Println(m)

	// m := map[string]int{"price": 30, "phone": 1}

	// k, ok := m["price"]

	// if ok {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("NOT OK")
	// }

	// fmt.Println(k)

	m1 := map[string]int{"price": 30, "phone": 1}
	m2 := map[string]int{"price": 30, "phone": 1}

	fmt.Println(maps.Equal(m1, m2))

}
