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
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

}

/*
[91 123 34 70 105 114 115 116 34 58 34 74 97 109 101 115 34 44 34 76 97 115
116 34 58 34 66 111 110 100 34 44 34 65 103 101 34 58 51 50 125 44 123 34 70
105 114 115 116 34 58 34 77 105 115 115 34 44 34 76 97 115 116 34 58 34 77 111
110 101 121 112 101 110 110 121 34 44 34 65 103 101 34 58 50 55 125 93]

[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]
*/
