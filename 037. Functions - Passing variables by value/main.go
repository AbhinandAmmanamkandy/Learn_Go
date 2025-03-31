package main

import "fmt"

func increment(x int) {
	x++
}

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
}

func main() {

	// Passing Variables by Value
	x := 5
	increment(x)
	fmt.Println(x)

	// Still prints 5
	// Since the increment function recieved a copy of x
	fmt.Println(monthlyBillIncrease(2, 89, 102))

}
