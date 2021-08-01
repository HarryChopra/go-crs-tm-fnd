package main

import "fmt"

func main() {
	x := 42
	y := 42.24

	fmt.Printf("var x is of type: %T, value: %v\n", x, x)
	fmt.Printf("var y is of type: %T, value: %v\n", y, y)

	s := "Hello world." // ASCII characters - 8bit/1 Byte, utf-8 - 32bit/ 4 Bytes
	bs := []byte(s)
	fmt.Printf("var s is of type: %T,\nvalue: %v\n", bs, bs)
	// [72 101 108 108 111 32 119 111 114 108 100 46]

	fmt.Printf("var s is of type: %T,\nvalue: %q\n", bs, bs) // string quoted representation
	// "Hello world."

	fmt.Printf("var s is of type: %T,\nvalue: %#v\n", bs, bs) // slice of bytes represented in Hex
	// []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e}

	z := "H"
	bsz := []byte(z)

	fmt.Printf("Binary representation: %b\n", bsz[0])
	fmt.Printf("Decimal representation: %d\n", bsz[0])
	fmt.Printf("Hexadecimal representation: %#x\n", bsz[0])

}
