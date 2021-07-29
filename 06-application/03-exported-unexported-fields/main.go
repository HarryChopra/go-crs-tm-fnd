package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	age   int // unexported field
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		age:   32,
	}
	fmt.Printf("%#v\n", p1) // main.person{First:"James", Last:"Bond", age:32}
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // {"First":"James","Last":"Bond"}
}
