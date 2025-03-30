package main

import "fmt"

func main() {

	// Constants - They can't use := short declaration - cannot be changed after declaration
	const pi = 3.14159

	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan: ", premiumPlanName)
	fmt.Println("plan: ", basicPlanName)

}
