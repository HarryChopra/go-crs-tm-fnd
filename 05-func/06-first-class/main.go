package main

import "fmt"

func main() {
	fmt.Println(calculate(add, 1, 2, 3, 4, 5))
	fmt.Println(calculate(multiply, 1, 2, 3, 4, 5))

}

func calculate(fn func(...int) int, x ...int) int {
	return fn(x...)
}

func add(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func multiply(x ...int) int {
	total := 1
	for _, v := range x {
		total *= v
	}
	return total
}

/*
15
120
*/
