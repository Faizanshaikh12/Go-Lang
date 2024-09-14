package main

import (
	"fmt"
)

// Most Importent Feature

// Communication between goroutines is done using channels.

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

// - Sending
// func processNum(nums chan int) {
// 	for num := range nums {
// 		fmt.Println("procesing number", num)
// 	}
// }

// Receiving
// func sum(result chan int, num1 int, num2 int) {
// 	sumResult := num1 + num2
// 	result <- sumResult
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("processing....")
// }

// send email
func emailSender(emailChan <-chan string, done chan<- bool) {

	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Email", email)
	}
}

func main() {

	// Buffer Channels Example
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	close(emailChan)
	<-done

	// done := make(chan bool)
	// go task(done)

	// <-done // block

	// result := make(chan int)

	// go sum(result, 4, 5)

	// fmt.Println(<-result)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// msg := <-messageChan

	// fmt.Println(msg)
}
