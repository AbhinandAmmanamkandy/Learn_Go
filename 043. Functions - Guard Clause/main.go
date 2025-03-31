package main

import "errors"

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// Here function returns early
		return 0, errors.New("Cannot divide by zero")
	}
	return dividend / divisor, nil
}

// Normal If else mess
// func getInsuranceAmount(status insuranceStatus) int {
// 	amount := 0
// 	if !status.hasInsurance(){
// 		amount = 1
// 	} else {
// 		if status.isTotaled() {
// 			amount = 10000
// 		}else {
// 			if status.isDented() {
// 				amount = 160
// 				if status.isBigDent() {
// 					amount = 270
// 				}
// 			}else{
// 				amount = 0
// 			}
// 		}
// 	}
// 	return amount
// }

// USING GUARD CLAUSE
// func getInsuranceAmount(status insuranceStatus) int{
// 	if !status.hasInsurance() {
// 		return 1
// 	}
// 	if status.isTotaled() {
// 		return 10000
// 	}
// 	if !status.isDented() {
// 		return 0
// 	}
// 	if status.isBigDent() {
// 		return 270
// 	}
// 	return 180
// }

func main() {
	// EARLY RETURNS
	// Go supports returning from a function early or continue through a loop

	// GUARD CLAUSES
	//
	// Guard Clauses leverage the ability to return early from a function
	// or continue through a loop to make nested conditionals one-dimentional
	// Instead of using if/else chains we just returns early from a function

	// Guard Clauses provide linear approach to logic trees
}
