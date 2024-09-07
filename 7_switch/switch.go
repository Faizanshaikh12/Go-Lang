package main

func main() {
	// simple switch

	// i := 3
	// switch i {
	// case 1:
	// 	println("One")
	// case 2:
	// 	println("Two")
	// case 3:
	// 	println("Three")
	// default:
	// 	println("Other")
	// }

	// multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("it's weekend")
	// default:
	// 	println("it's workday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("its an integer")
		case string:
			println("its a string")
		case bool:
			println("its a boolean")
		default:
			println("other", t)
		}
	}

	whoAmI(50.05)
}
