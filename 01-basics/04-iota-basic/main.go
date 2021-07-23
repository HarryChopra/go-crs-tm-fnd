package main

import "fmt"

const (
	// start from 0
	a = iota
	b
	c
)

const (
	// start from 1
	x = iota + 1
	y
	z
)

const (
	// skip a value using blank identifier
	c1 = iota + 1
	_
	c2
	c3
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
	fmt.Println(c1, c2, c3)
}
