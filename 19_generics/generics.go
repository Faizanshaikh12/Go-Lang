package main

import "fmt"

// LIFO - Last In First Out
type stack[T any] struct {
	elements []T
}

func printSlice[T comparable](nums []T) {
	for _, item := range nums {
		fmt.Println(item)
	}
}

func main() {
	myStack := stack[string]{
		elements: []string{"Faizan"},
	}

	fmt.Println(myStack)

	nums := []int{1, 2, 3}
	values := []string{"Faizan", "Shaikh"}
	printSlice(nums)
	printSlice(values)
}
