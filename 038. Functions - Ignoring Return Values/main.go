package main

import "fmt"

func getPoint() (x int, y int) {
	return 3, 4
}

func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier) 	// WE IGNORED THE THIRD STRING RETURN TYPE USING UNDERSCORE
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

func main() {

	// IGNORING RETURN VALUES

	// EVEN THOUGH getPoint() RETURNS TWO VALUES
	x, _ := getPoint() // WE IGNORED SECOND RETURN VALUE USING UNDERSCORE
	fmt.Println(x)

	fmt.Println(getProductMessage("premium"))

}
