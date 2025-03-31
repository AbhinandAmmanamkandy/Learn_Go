package main

import "fmt"

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	finalCost := calculateBaseBill(costPerMessage, numMessages)
	discount := calculateDiscountRate(numMessages)
	finalCost -= finalCost * discount
	return finalCost
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	}
	if messagesSent > 1000 {
		return 0.1
	}
	return 0
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main() {
	fmt.Println(calculateFinalBill(1.0, 1428))
}
