package main

import "fmt"

func main() {

	// Type Inference - Assigning type upon declaration
	// %T is used to print type of variable in Go lang

	// penniesPerText := 2 // Is int
	penniesPerText := 2.0 // Is float

	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

}
