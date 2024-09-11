package main

import "fmt"

// by params
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In Change Num", num)
// }

// by memory reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In Change Num", *num)
}

func changeSlice(nums *[]int) {
	*nums = append(*nums, 4)
	fmt.Println("In Change Slice", *nums)
}

func main() {
	num := 1

	// changeNum(num)
	changeNum(&num)

	fmt.Println("After Change Num", num)

	nums := []int{1, 2, 3}

	changeSlice(&nums) // & Gettng Variable Memory Reference

	fmt.Println("After Change Slice", nums)
}
