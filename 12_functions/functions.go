package main

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "Faizan", "Shaikh", true
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// result := add(3, 5)
	// fmt.Println(result)

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn := processIt()
	fn(6)

}
