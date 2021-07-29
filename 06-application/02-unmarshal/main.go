package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	rawJSON := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	people := []person{}
	if err := json.Unmarshal([]byte(rawJSON), &people); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", people)
	fmt.Println(people[0].First)
}

/*
[]main.person{main.person{First:"James", Last:"Bond", Age:32},
main.person{First:"Miss", Last:"Moneypenny", Age:27}}

James
*/
