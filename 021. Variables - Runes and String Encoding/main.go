package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Emoji has a byte length of 4

	// name := "boots"
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
