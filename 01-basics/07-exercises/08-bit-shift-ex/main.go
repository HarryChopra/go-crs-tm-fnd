package main

import "fmt"

func main() {
	x := 40
	fmt.Printf("var x in binary representation: %b\n", x)
	fmt.Printf("var x in decimal representation: %d\n", x)
	fmt.Printf("var x in hex representation: %#x\n", x)

	x = x << 1
	fmt.Println("After the bit shift:")
	fmt.Printf("var x in binary representation: %b\n", x)
	fmt.Printf("var x in decimal representation: %d\n", x)
	fmt.Printf("var x in hex representation: %#x\n", x)

	x = x << 1
	fmt.Println("After the bit shift:")
	fmt.Printf("var x in binary representation: %b\n", x)
	fmt.Printf("var x in decimal representation: %d\n", x)
	fmt.Printf("var x in hex representation: %#x\n", x)
}

/*
Output:

var x in binary representation: 101000
var x in decimal representation: 40
var x in hex representation: 0x28

After the bit shift:
var x in binary representation: 1010000
var x in decimal representation: 80
var x in hex representation: 0x50

After the bit shift:
var x in binary representation: 10100000
var x in decimal representation: 160
var x in hex representation: 0xa0
*/
