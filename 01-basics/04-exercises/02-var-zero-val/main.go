package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	fmt.Printf("%T\t%d\n", x, x)
	fmt.Printf("%T\t%s\n", y, y)
	fmt.Printf("%T\t%t\n", z, z)
}
