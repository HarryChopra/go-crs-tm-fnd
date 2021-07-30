package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Last  string
	Age   int
}

type byFirst []person
type byLast []person
type byAge []person

// Type byFirst implements sort.Sort interface
func (f byFirst) Len() int           { return len(f) }
func (f byFirst) Less(x, y int) bool { return f[x].First < f[y].First }
func (f byFirst) Swap(x, y int)      { f[x], f[y] = f[y], f[x] }

// Type byLast implements sort.Sort interface
func (f byLast) Len() int           { return len(f) }
func (f byLast) Less(x, y int) bool { return f[x].Last < f[y].Last }
func (f byLast) Swap(x, y int)      { f[x], f[y] = f[y], f[x] }

// Type byAge implements sort.Sort interface
func (f byAge) Len() int           { return len(f) }
func (f byAge) Less(x, y int) bool { return f[x].Age < f[y].Age }
func (f byAge) Swap(x, y int)      { f[x], f[y] = f[y], f[x] }

func main() {
	people := []person{
		{First: "James", Last: "Bond", Age: 32},
		{First: "Miss", Last: "Moneypenny", Age: 27},
		{First: "Ian", Last: "Fleming", Age: 56},
	}
	sort.Sort(byFirst(people))
	fmt.Println("Sorted by First Name:\n", people)

	sort.Sort(byLast(people))
	fmt.Println("Sorted by Last Name:\n", people)

	sort.Sort(byAge(people))
	fmt.Println("Sorted by Age:\n", people)
}

/*
Sorted by First Name:
 [{Ian Fleming 56} {James Bond 32} {Miss Moneypenny 27}]

Sorted by Last Name:
 [{James Bond 32} {Ian Fleming 56} {Miss Moneypenny 27}]

Sorted by Age:
 [{Miss Moneypenny 27} {James Bond 32} {Ian Fleming 56}]
*/
