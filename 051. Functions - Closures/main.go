package main

import "fmt"

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func adder() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {

	// CLOSURE
	// A closure is a function that references variables from outside its own function body
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))

	add := adder()
	add(1)
	add(2)
	fmt.Println(add(3))

}
