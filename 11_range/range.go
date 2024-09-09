package main

import "fmt"

// iterating over data structures
func main() {
	fmt.Println("Range")

	// nums := []int{6, 7, 8}

	// sum := 0

	// for idx, num := range nums {
	// 	sum = sum + num
	// 	fmt.Println(idx, num)
	// }

	// fmt.Println(sum)

	// m := map[string]string{"fname": "Faizan", "lname": "Shaikh"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// unicode code point rune
	// starting byte of rune
	for i, c := range "Faizan" {
		fmt.Println(i, c, string(c))
	}

}
