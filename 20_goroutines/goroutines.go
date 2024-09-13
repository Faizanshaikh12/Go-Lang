package main

import (
	"fmt"
	"sync"
)

// Go most powerful feature

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		// go keyword - Multi threading
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	time.Sleep(time.Second * 1)
// }

// Wait Group - Without using time sleep completed function runing
func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing task", i)
}

func main() {
	// receive wait group
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
