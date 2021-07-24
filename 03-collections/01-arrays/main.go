package main

import "fmt"

func main() {
	var arr1 [3]int

	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	fmt.Printf("var arr1 is of type: %T, has val:%#v\n", arr1, arr1)

	// var arr2 [3]string = [3]string{"No time to die", "Spectre", "Skyfall"}
	// arr2 := [3]string{"No time to die", "Spectre", "Skyfall"} // composite literal
	arr2 := [...]string{"No time to die", "Spectre", "Skyfall"} // array literal size using three dot notation
	fmt.Printf("var arr2 is of type: %T, has val:%#v\n", arr2, arr2)
}

/*
output:

var arr1 is of type: [3]int, has val:[3]int{10, 20, 30}
var arr2 is of type: [3]string, has val:[3]string{"No time to die", "Spectre", "Skyfall"}
*/
