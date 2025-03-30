package main

import "fmt"

func main() {
	// Type sizes

	// Whole numbers (No Decimal)
	// int int8 int16 int32 int64

	// Positive Whole Numbers (No Decimal)
	// unit uint8 uint16 uint32 uint64 uintptr

	// Signed Decimal Numbers
	// float32 float64

	// Imaginary Numbers
	// complex64 complex128

	// The size 8, 16, 32, 64, 128 etc represents bits

	// Converting between types
	temperatureFloat := 88.26
	temperatureInt := int64(temperatureFloat)
	fmt.Println(temperatureInt)

	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
