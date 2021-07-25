package main

import "fmt"

func main() {
	// If a one off struct is needed, use an anonymous struct and prevent code pollution
	p1 := struct {
		First, Last string
		Age         int
	}{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	fmt.Printf("%+v\n", p1)
}
