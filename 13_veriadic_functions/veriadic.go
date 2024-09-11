package main

import "fmt"

func sum(nums ...int) int {

	total := 0

	for _, num := range nums {
		total = num + total
	}

	return total

}

func main() {
	nums := []int{1, 2, 3}
	result := sum(nums...)
	fmt.Println(result)
}
