package main

import "fmt"

func main() {
	// This is a single line comment

	/*
		This is a multi-line comment
		neither of these comments will execute
		as code
	*/

	/*
		We are increasing the maximum message length from 140 to 280 characters
		very reluctantly, I might add
		Users actually want to write more than 140 characters
	*/
	maxMessageLength := 140
	newMaxMessageLength := 280
	fmt.Println("Textio is increasing the maximum message length from", maxMessageLength, "to", newMaxMessageLength)
}
