package main

import "fmt"

type Person struct {
	Name string
}

func ChangeName1(p Person) {
	p.Name = "Updated Name"
}

func ChangeName2(p *Person) {
	p.Name = "JB"
	// (*p).Name = "JB"
}

func main() {
	p := Person{
		Name: "James Bond",
	}

	ChangeName1(p) // pass value
	fmt.Println(p) // James Bond (Unchanged)

	ChangeName2(&p) // Pass reference
	fmt.Println(p)  // JB
}
