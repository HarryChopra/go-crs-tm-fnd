package main

import "fmt"

type name string

var x name

func main() {
	fmt.Printf("The type of x: %T\n", x)
	fmt.Printf("The default value of x: %#v\n", x)

	x = "Steve"
	fmt.Printf("The assigned value of x: %s\n", x)
}
