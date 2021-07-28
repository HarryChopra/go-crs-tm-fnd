package main

import "fmt"

func main() {
	fmt.Println("factorial of 0 is:", factorial(0))
	fmt.Println("factorial of 1 is:", factorial(1))
	fmt.Println("factorial of 2 is:", factorial(2))
	fmt.Println("factorial of 3 is:", factorial(3))
	fmt.Println("factorial of 4 is:", factorial(4))
}

func factorial(x int) int {
	if x > 1 {
		return factorial(x-1) * x
	}
	return 1
}

/*
factorial of 0 is: 1
factorial of 1 is: 1
factorial of 2 is: 2
factorial of 3 is: 6
factorial of 4 is: 24
*/
