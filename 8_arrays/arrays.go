package main

import "fmt"

// numbered sequence of specific length
func main() {
	// zero values int->0, string->"", bool->false
	// var nums [4]int

	// nums[0] = 1

	// arrays length
	// fmt.Println(len(nums))

	// print arrays
	// fmt.Println(nums)

	// var vals [4]bool
	// vals[0] = true
	// fmt.Println(vals)

	// to declare it in single line
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	// 2d arrays
	arr := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(arr)

	// - Fixed size, that is predictable
	// - Memory optimization
	// - Contant time access
}
