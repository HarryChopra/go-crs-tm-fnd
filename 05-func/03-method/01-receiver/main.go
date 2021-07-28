package main

import "fmt"

type Person struct {
	First, Last string
}

type SecretAgent struct {
	Person
	LTK bool
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			First: "James",
			Last:  "Bond",
		},
		LTK: true,
	}

	sa1.speak()

}

// method receiver
func (p Person) speak() {
	fmt.Printf("Hi, I am %s %s.\n", p.First, p.Last)
}

/*
Hi, I am James Bond.
*/
