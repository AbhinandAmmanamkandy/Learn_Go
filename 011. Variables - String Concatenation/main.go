package main

import "fmt"

func main() {
	// Go is static typed language - variable types are known before code runs
	// Means editor and compiler displays error before the code even run

	// Dynamically typed languages - JavaScript and Python

	// String Concatenation
	// Two strings can be concatenated using + operator
	// But you cannot concatenate string with int or float64

	var username string = "Ghost Reborn"
	var password string = "12345"
	fmt.Println("Authorization: Basic", username+":"+password)

}
