package main

import "fmt"

func getCoords() (x, y int) {
	return x, y // This is explicit
}

func getCoord() (x, y int) {
	return 5, 6 // This is explicit, x and y are NOT returned
}

func getCoo() (x, y int) {
	return
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
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

	// EXPICIT RETURNS - Function explicitly returns the value
	fmt.Println(yearsUntilEvents(22))
}
