package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("%U '%c'\n", i, i)
	}
}

/*
Output:

U+0041'A'
U+0042'B'
U+0043'C'
...
U+005A'Z'
*/
