package main

import "fmt"

type Person struct {
	First, Last string
	Age         int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   26,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.Last, p2.Last, p1.Age)

}
