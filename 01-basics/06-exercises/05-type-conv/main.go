package main

import "fmt"

type leader string

var x leader = "MLK"

func main() {
	fmt.Printf("The type and value of x is: %T, %s\n", x, x)

	y := "Winston Churchill"
	x = leader(y)

	fmt.Printf("The type and value of x is: %T, %s\n", x, x)
}
