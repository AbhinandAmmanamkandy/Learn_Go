package main

import "fmt"

func main() {

	// Same line declarations
	mileage, company := 80276, "Toyota"
	fmt.Printf("mileage: %v, company: %v\n", mileage, company)

	averageOpenRate := 0.23
	displayMessage := "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)

}
