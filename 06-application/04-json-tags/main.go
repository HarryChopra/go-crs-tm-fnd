package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"firstName"`
	Last  string `json:"lastName"`
	Age   int    `json:"age,omitempty"` // Omit if zero value or nil
	LTK   bool   `json:"-"`             // Unexported field
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		LTK:   true,
	}

	p2 := person{
		First: "Miss",
		LTK:   false,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	// [{"firstName":"James","lastName":"Bond","age":32},{"firstName":"Miss","lastName":""}]
	// p2's "Age" and "Last" field are empty, yet, "Last" is exported to json
}
