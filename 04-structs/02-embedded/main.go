package main

import "fmt"

type Person struct {
	First, Last string
	Age         int
}

type SecretAgent struct {
	Person
	Age int // Name collision.  Top level type's field value takes precedence.
	LTK bool
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   32,
		},
		LTK: true,
	}

	fmt.Printf("%+v\n", sa1)
	fmt.Println(sa1.Last)       // Bond (Inner type field is promoted)
	fmt.Println(sa1.Person.Age) // 32 (Because of name collision, Inner type isnt promoted)
	fmt.Println(sa1.Age)        // 0

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   26,
	}

	sa2 := SecretAgent{
		Person: p2,
		Age:    24, // This assignment takes precedence
		LTK:    false,
	}

	fmt.Printf("%+v\n", sa2)
	fmt.Println(sa2.Age) // 24
}
