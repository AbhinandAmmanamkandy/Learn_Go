package main

import "fmt"

// Here, func sub(x int, y int) int is known as the "function signature".
func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {

	fmt.Println(sub(3, 2))
	test("Ghost", " Reborn")

}
