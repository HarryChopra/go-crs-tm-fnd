package main

import "fmt"

type Human interface {
	Speak()
}

type Person struct {
	Name string
	Age  string
}

func (p Person) Speak() {
}

type SecretAgent struct {
	Person
	Name string
	LTK  bool
}

func (sa SecretAgent) Speak() {
}

func Greet(h Human) {
	// switch h.(type) {
	switch h := h.(type) {
	case Person:
		// fmt.Println("Person: Hi, I'm", h.(Person).Name) // Assert Type
		fmt.Println("Person: Hi, I'm", h.Name)
	case SecretAgent:
		fmt.Println("Secret Agent: Hi, I'm", h.Person.Name)
	}
}

func main() {
	sa := SecretAgent{
		Person: Person{
			Name: "James Bond",
			Age:  "32",
		},
		Name: "JB",
		LTK:  true,
	}
	Greet(sa)

	p := Person{
		Name: "Miss Moneypenny",
		Age:  "27",
	}
	Greet(p)
}
