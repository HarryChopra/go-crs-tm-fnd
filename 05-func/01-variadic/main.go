package main

import "fmt"

func main() {
	result := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("sum total: ", result)

	sl := []int{10, 20, 30, 40}
	result = sum(sl...) // unfurling slice elements
	fmt.Println("new sum total: ", result)
}

// variadic params
func sum(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}
