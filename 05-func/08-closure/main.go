package main

import "fmt"

func main() {
	i := incrementer()
	fmt.Println("i:", i())
	fmt.Println("i:", i())
	fmt.Println("i:", i())
	fmt.Println("i:", i())
	fmt.Println("i:", i())

	j := incrementer()
	fmt.Println("j:", j())
	fmt.Println("j:", j())
	fmt.Println("j:", j())
}

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

/*
i: 1
i: 2
i: 3
i: 4
i: 5

j: 1
j: 2
j: 3
*/
