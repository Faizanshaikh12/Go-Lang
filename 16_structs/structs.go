package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	order1 := order{
		id:        "1",
		amount:    100,
		status:    "Paid",
		createdAt: time.Now(),
		customer: customer{
			name:  "Faizan",
			phone: "9664979557",
		},
	}

	return &order1
}

func main() {

	order1 := order{
		id:        "1",
		amount:    100,
		status:    "Paid",
		createdAt: time.Now(),
	}

	order2 := order{
		id:        "2",
		amount:    200,
		status:    "Un Paid",
		createdAt: time.Now(),
	}

	fmt.Println("order1", order1)

	order2.changeStatus("Confirmed")

	fmt.Println("order2", order2)
	fmt.Println("order2 amount", order2.getAmount())

	// Create New Order
	order3 := newOrder("3", 10.10, "Paid")
	fmt.Println("order3", order3)

	// passing default value
	lang := struct {
		name   string
		isGood bool
	}{"Hindi", true}
	fmt.Println("language", lang)

}
