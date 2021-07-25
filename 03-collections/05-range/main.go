package main

import "fmt"

func main() {
	s := [][]string{
		{"James", "Bond", "Shaken not stirred"},
		{"Miss", "Moneypenny", "Hellooo, James"},
	}

	for i1, v1 := range s {
		fmt.Println(i1)
		for i2, v2 := range v1 {
			fmt.Printf("\t%d: %s\n", i2, v2)
		}
	}

	m := map[string][]string{
		"bond":       {"bond_james", "Shaken, not stirred", "Martinis", "Women"},
		"moneypenny": {"moneypenny_miss", "James Bond", "Literature", "Computer Science"},
		"no":         {"no_dr", "Being evil", "Ice cream", "Sunsets"},
	}

	for k1, v1 := range m {
		fmt.Println(k1)
		for i2, v2 := range v1 {
			fmt.Printf("\t%d: %s\n", i2, v2)
		}
	}

}
