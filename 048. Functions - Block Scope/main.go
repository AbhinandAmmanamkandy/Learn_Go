package main

import "fmt"

func splitEmail(email string) (string, string) {
	// This is wrong since it's scope is inside this brackets
	// {
	// 	username, domain := "", ""
	// }

	// This is right
	{
		username, domain := "", ""
	}
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}


// scoped to the entire "main" package (basically global)
var age = 19

func main() {

	// Block Scope
	//
	// Variables declared inside a block are only accessible within that block
	// and its nested blocks

	// Blocks are defined by curly braces {}
	// New blocks are created for
	//
	// functions
	// loops
	// if statements
	// switch statements
	// select statements
	// explicit blocks

	// scoped to the "main" function
	name := "Jon Snow"

	for i := 0; i < 5; i++ {
		// scoped to the "for" body
		email := "snow@winterfell.net"
		fmt.Println(email)
	}
	fmt.Println(name)

	{
		age := 23
		// This is okay
		fmt.Println(age)
	}

	// This is not okay
	fmt.Println(age) // This works since we have a global variable with value 19

}
