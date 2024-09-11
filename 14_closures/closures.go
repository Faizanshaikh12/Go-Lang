package main

import "fmt"

func counter() func() int {
	var count int = 0

	// Jab bhi apka stack empty ho jayega fir bhi apko function me outer variable milege apko return function me
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
}
