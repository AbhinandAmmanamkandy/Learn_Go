package main

import "fmt"

func main() {

	// Basic Variables

	// bool
	// string
	// int
	// float64
	// byte

	// Declaring a variable the sad way
	var mySkillIssues int
	mySkillIssues = 42
	fmt.Println(mySkillIssues)

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

}
