package main

import "errors"

// This is easier to understand than
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot diivde by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// This
func calc(a, b int) (int, int, error){
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	mul := a * b
	div := a / b
	return mul, div, nil
}

func main() {

	// Named return paramters are great for documenting a function
	// When might you use naked returns?
	// For small functions

}
