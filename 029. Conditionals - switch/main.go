package main

import "fmt"

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linux Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "mac":
		creator = "A steve"
	default:
		creator = "Unknown"
	}
	return creator
}

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func main() {
	// switch - compare value against multiple options
	creator := getCreator("linux")
	fmt.Println(creator)

	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}
