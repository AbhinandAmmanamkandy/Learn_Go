package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := 0.2
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}
