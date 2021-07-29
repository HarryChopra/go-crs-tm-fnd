package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type colorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type total struct {
	A *colorGroup `json:",omitempty"`
	B string      `json:",omitempty"`
}

func main() {
	groupWithNilA := total{
		B: "lorem",
	}
	bs, err := json.Marshal(groupWithNilA)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(string(bs)) // {"B":"lorem"}

	groupWithA := total{
		A: &colorGroup{},
		B: "ipsum",
	}
	bs, err = json.Marshal(groupWithA)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(string(bs)) // {"A":{"ID":0,"Name":"","Colors":null},"B":"ipsum"}
}

/*
Struct values encode as JSON objects. Each exported struct field becomes a member of the object unless

- the field's tag is "-", or
- the field is empty and its tag specifies the "omitempty" option.
- The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero.
*/
