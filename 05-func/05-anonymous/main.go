package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println("The sum is:", a+b)
	}(2, 3)

	result := func(a, b int) int {
		return a * b
	}(2, 3)

	fmt.Println("The product is:", result)
}

/*
The sum is: 5
The product is: 6
*/
