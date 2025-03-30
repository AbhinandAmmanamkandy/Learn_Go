package main

import "fmt"

func main() {
	// Formatting strings in Go

	// fmt.Printf - Prints a formatted string to standard out
	// fmt.Sprintf() - Returns the formatted string

	// %v - Go syntax representation of a value

	// String

	s := fmt.Sprintf("I am %v years old", 10)
	fmt.Println(s)

	s = fmt.Sprintf("I am %v years old", "way too many")
	fmt.Println(s)

	// Integer
	i := fmt.Sprintf("I am %d years old", 10)
	fmt.Println(i)

	// Float
	f := fmt.Sprintf("I am %.2f years old", 10.523)
	fmt.Println(f)

	const name = "Ghost Reborn"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
	fmt.Print(msg)
}
