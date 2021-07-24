package main

import "fmt"

func main() {
	for i := 2022; i > 1990; i-- {
		fmt.Println(i)
	}

	i := 1990
	for {
		fmt.Println(i)
		i++
		if i > 2022 {
			break
		}
	}

	for i := 0; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
