package main

import "fmt"

func conversions(convertor func(int) int, x, y, z int) (int, int, int) {
	convertedX := convertor(x)
	convertedY := convertor(y)
	convertedZ := convertor(z)
	return convertedX, convertedY, convertedZ
}

func double(a int) int {
	return a + a
}

func getCost(message string) int{
	return len(message) * 2
}

func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return len(s) * 2
	}, intro)
	printCostReport(func (message string) int{
		return len(message) * 3
	}, body)
	printCostReport(func (message string) int {
		return len(message) * 4
	}, outro)
}

func printCostReport(costCalculator func(string) int, message string){
	cost := costCalculator(message)
	fmt.Printf("Message: %s Cose: %v cents", message, cost)
	fmt.Println()
}

func main() {

	// Anonymous functions
	//
	// Anonymous functions are true to form in that they have no name.
	// They are useful when definig a function that will only be used once to create a quick closure

	// Using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	fmt.Println(newX, newY, newZ)

	newX, newY, newZ = conversions(func (a int) int {
		return a + a
	}, 1, 2, 3)
	fmt.Println(newX, newY, newZ)

	printReports("Intro", "Body", "Outro")

}
