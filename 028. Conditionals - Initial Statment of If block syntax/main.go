package main

import "fmt"

func main() {

	// Invalid Syntax

	// if(x:=5;x>3) {
	// 	fmt.Println("x is greater than 3")
	// }

	// Valid Syntax

	if x := 5; x > 3 {
		fmt.Println("x is greater than 3")
	}

}
