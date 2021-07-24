package main

import "fmt"

const (
	a = 10
	b = "Hello World"
	c = true
)

const (
	x int    = 20
	y string = "Foo"
	z bool   = false
)

func main() {
	fmt.Printf("const a, b and c are of type: %T, %T and %T\n", a, b, c)
	fmt.Printf("const x, y and z are of type: %T, %T and %T\n", x, y, z)
}
