package main

func main() {
	age := 16

	if age >= 18 {
		println("Person is an adult")
	} else if age >= 12 {
		println("Person is teenager")
	} else {
		println("Person is kid")
	}

	// go does not have ternary, you will have to use normal if else.
}
