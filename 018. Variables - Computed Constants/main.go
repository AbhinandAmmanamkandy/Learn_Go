package main

import (
	"fmt"
)

func main() {

	// Computed Constants
	// Constants must be known at compile time. They are usually declared with a static value
	const myInt = 15

	// Example for Computed Constants
	const firstName = "Ghost"
	const lastName = "Reborn"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)

	// You cannot declare a constant that can only be computed on runtime
	// const currentTime = time.Now()

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour
	fmt.Println("number of seconds in an hour:", secondsInHour)

}
