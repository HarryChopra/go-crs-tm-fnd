package main

import "fmt"

func main() {
	sl1 := []int{2, 3, 5, 7, 11, 15}
	fmt.Printf("len=%d, cap=%d, %v\n", len(sl1), cap(sl1), sl1)

	// using make
	sl2 := make([]string, 0)
	printSlice(sl2)

	// Specifying the capacity to avoid unnecessary allocation:
	// As the size of the slice grows, the runtime doubles its capacity,
	// every time it exceeds its current capacity. It has to allocate a new slice
	// with double capacity and copy over all of the elements from the slice.
	// If we know how big the slice will be from the start, we can avoid both:
	// copy operations and memory allocations.
	// NOTE: After 1024 elements, the runtime starts increasing capacity in 25% increments.
	sl3 := make([]string, 0, 10)
	printSlice(sl3)

	sl3 = append(sl3, "a", "b", "c", "d", "e", "f", "g", "h", "i")
	printSlice(sl3)

	sl3 = append(sl3, "j", "k", "l")
	printSlice(sl3)

}

func printSlice(sl []string) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(sl), cap(sl), sl)
}

/*
Output:

len=6, cap=6, [2 3 5 7 11 15]
len=0, cap=0, []
len=0, cap=10, []
len=9, cap=10, [a b c d e f g h i]
len=12, cap=20, [a b c d e f g h i j k l]
*/
