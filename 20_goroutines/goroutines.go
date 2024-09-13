package main

import (
	"fmt"
	"time"
)

// Go most powerful feature

func main() {
	for i := 0; i <= 10; i++ {
		// go keyword - Multi threading
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}
