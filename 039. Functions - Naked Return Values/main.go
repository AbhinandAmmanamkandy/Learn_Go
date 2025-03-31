package main

import "fmt"

func getCoord() (x, y int) {
	// x and y are initialized with zero values
	return // automatically returns x and y
}

// ABOVE IS SAME AS BELOW
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return
}

func main() {

	// NAKED RETURN
	// A return statement without arguments returns the named return values
	// This is known as a "naked" return
	// This should only be used for short functions
	// They can harm readability in longer functions

	fmt.Println(yearsUntilEvents(23))

}
