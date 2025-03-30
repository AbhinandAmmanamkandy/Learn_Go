package main

import "fmt"

func main() {
	fname := "Ghost"
	lname := "Reborn"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %v %v, Age: %v, Rate: %v, Is Subscribed: %v, Message: %v", fname, lname, age, messageRate, isSubscribed, message)
	fmt.Println(userLog)
}
