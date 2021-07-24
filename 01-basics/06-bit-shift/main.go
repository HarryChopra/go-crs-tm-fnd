package main

import "fmt"

const (
	KB = 1 << ((iota + 1) * 10)
	MB
	GB
)

func main() {
	x := 4
	fmt.Printf("Decimal representaion of x: %d\n", x) // 4
	fmt.Printf("Binary representaion of x: %b\n", x)  // 100

	y := x << 1
	fmt.Printf("Decimal representaion of y: %d\n", y) // 8
	fmt.Printf("Binary representaion of y: %b\n", y)  // 1000

	kb := 1024
	mb := kb * kb
	gb := mb * kb

	fmt.Printf("Binary representation:\nkb: %b\nmb: %b\ngb: %b\n", kb, mb, gb)
	// kb: 10000000000
	// mb: 100000000000000000000
	// gb: 1000000000000000000000000000000

	// Implement the bit shifting with iota (const declaration above)
	fmt.Printf("Binary representation:\nKB: %b\nMB: %b\nGB: %b\n", KB, MB, GB)
	// KB: 10000000000
	// MB: 100000000000000000000
	// GB: 1000000000000000000000000000000
}
