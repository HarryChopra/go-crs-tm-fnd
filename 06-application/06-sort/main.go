package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 9, 6, 7, 8, 1, 3, 2, 5, 4}
	letters := []string{"a", "z", "b", "y", "c", "x"}

	// sort the slice
	sort.Ints(nums)
	fmt.Println(nums)

	// reverse the slice
	sort.SliceStable(nums, func(i, j int) bool {
		return i > j
	})
	fmt.Println(nums)

	sort.Strings(letters)
	fmt.Println(letters)
}
