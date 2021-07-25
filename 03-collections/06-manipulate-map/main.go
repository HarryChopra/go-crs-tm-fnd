package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond":       {"bond_james", "Shaken, not stirred", "Martinis", "Women"},
		"moneypenny": {"moneypenny_miss", "James Bond", "Literature", "Computer Science"},
		"no":         {"no_dr", "Being evil", "Ice cream", "Sunsets"},
	}

	// add a record
	m["fleming"] = []string{"fleming_ian", "cigars", "espionage"}
	printMap(m)

	// remove a record
	fmt.Println("deleting a record:")
	delete(m, "no")
	printMap(m)

}

func printMap(m map[string][]string) {
	for k1, v1 := range m {
		fmt.Println(k1)
		for i2, v2 := range v1 {
			fmt.Printf("\t%d: %s\n", i2, v2)
		}
	}
}
