package main

import "fmt"

func main() {
	names := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Printf("%T\n", names) // [4]string

	// 1. Slicing an array returns a slice
	a := names[0:2]
	b := names[1:3]

	fmt.Printf("%T\t%T\n", a, b) // []string        []string
	fmt.Println(a, b)            // [John Paul] [Paul George]

	// A slice literal is like an array literal without the length.

	// This is an array literal:
	// [3]bool{true, true, false}

	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}

	// 2. Slice is like a reference to an array
	a[0] = "Peter"
	fmt.Println(a)     // [Peter Paul]
	fmt.Println(names) // [Peter Paul George Ringo]

	// 3. Zero value of a slice is nil
	var sl1 []int                                                     // variable declaration
	fmt.Println(sl1 == nil)                                           // true
	fmt.Printf("len=%d, cap=%d, sl1 = %v\n", len(sl1), cap(sl1), sl1) // len=0, cap=0, sl1 = []

	sl2 := []int{}          // variable declaration and assignment
	fmt.Println(sl2 == nil) // false

	sl3 := make([]int, 0)   // variable declaration and assignment
	fmt.Println(sl3 == nil) // false

}
