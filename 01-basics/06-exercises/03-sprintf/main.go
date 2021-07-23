package main

import "fmt"

var (
	x = 42
	y = "James Bond"
	z = true
)

func main() {
	str := fmt.Sprintf("the values of x, y & z are: %d, %s, %t", x, y, z)
	fmt.Println(str)
}
