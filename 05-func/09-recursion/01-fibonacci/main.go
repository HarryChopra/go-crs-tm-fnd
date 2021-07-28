package main

import "fmt"

func main() {

	fmt.Println("Fibonacci of 1 is:", fib(1))
	fmt.Println("Fibonacci of 2 is:", fib(2))
	fmt.Println("Fibonacci of 2 is:", fib(3))
	fmt.Println("Fibonacci of 8 is:", fib(8))
	fmt.Println("Fibonacci of 13 is:", fib(13))
}

func fib(x int) int {
	if x > 2 {
		return fib(x-1) + fib(x-2)
	}
	return 1
}

/*
Fibonacci of 1 is: 1
Fibonacci of 2 is: 1
Fibonacci of 2 is: 2
Fibonacci of 8 is: 21
Fibonacci of 13 is: 233
*/
