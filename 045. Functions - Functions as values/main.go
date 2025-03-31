package main

import "fmt"

// below are 2 simple functions
func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithematic func(int, int) int) int {
	firstResult := arithematic(a, b)
	secondResult := arithematic(firstResult, c)
	return secondResult
}

func reformat(message string, formatter func(string) string) string {
	message = "TEXTIO: " + formatter(formatter(formatter(message)))
	return message
}

func main() {

	// Functions as values
	//
	// Go supports first class and higher order functions
	// Means "functions can be assigned as values"
	// Functions are just another type just like int, string, bools

	sum := aggregate(2, 3, 4, add)
	fmt.Println(sum)

	product := aggregate(2, 3, 4, mul)
	fmt.Println(product)

}
