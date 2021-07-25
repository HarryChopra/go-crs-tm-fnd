package main

import "fmt"

type Person struct {
	First, Last      string
	FavoriteFlavours []string
}

func main() {
	p1 := Person{
		First:            "James",
		Last:             "Bond",
		FavoriteFlavours: []string{"Chocolate", "Martini", "Rum and Coke"},
	}
	printPerson(p1)

	p2 := Person{
		First:            "Miss",
		Last:             "MoneyPenny",
		FavoriteFlavours: []string{"Strawberry", "Vanilla", "Capuccino"},
	}
	printPerson(p2)
}

func printPerson(p Person) {
	fmt.Printf("%s %s. Favorite Flavours: ", p.First, p.Last)
	for _, fl := range p.FavoriteFlavours {
		fmt.Printf("%s,", fl)
	}
	fmt.Println()
}

/*
James Bond, Favorite Flavours: Chocolate,Martini,Rum and Coke,
Miss MoneyPenny, Favorite Flavours: Strawberry,Vanilla,Capuccino,
*/
