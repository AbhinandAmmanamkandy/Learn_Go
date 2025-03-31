package main

import "fmt"

func getUserName(dstName, srcName string) (username string, err error) {
	// Open a connection to the database
	conn, _ := db.Open(srcName)

	// Close the connection *anywhere* the getUserName function returns
	defer conn.Close()

	username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil

	// In the above example conn.Close() is not called on line 8
	// But called on line 13 or line 17

}

func bootup() {

	defer fmt.Println("TEXTIO BOOTUP DONE") // Prints TEXTIO BOOTUP DONE after bootup function is completed

	ok := connectToDB()
	if !ok {
		return
	}

	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All Systems ready!")

}

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed!")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("=======================================")
}

func main() {

	// defer keyword is a fairly unique feature for Go
	// Deferred functions are typically used to clean up resources
	// that are no longer being used
	// Often to close database connections, file handlers

	test(true, true)

}
