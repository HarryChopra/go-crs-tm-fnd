package main

import "fmt"

func main() {
	// Assign function to a variable
	sum := printSum
	sum(5, 10)

	product := printProduct
	product(5, 10)

	// Return a function
	sum = displaySum()
	sum(10, 10)
}

func printSum(a, b int) {
	fmt.Println(a + b)
}

func printProduct(a, b int) {
	fmt.Println(a * b)
}

func displaySum() func(int, int) {
	return printSum
}

/*
15
50
20
*/
