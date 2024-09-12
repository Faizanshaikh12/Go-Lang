package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// open close principle
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe gateway for testing purpose", amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay gateway for testing purpose", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal gateway for testing purpose", amount)
}

func main() {

	// stripeGw := stripe{}
	// razorpayGw := razorpay{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}

	newPayment.makePayment(100)
}
