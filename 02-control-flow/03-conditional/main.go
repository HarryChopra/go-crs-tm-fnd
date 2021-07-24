package main

import "fmt"

func main() {
	x := "James Bond"

	if x == "Miss Moneypenny" {
		fmt.Println("Hellooo James")
	} else if x == "James Bond" {
		fmt.Println("Bond. James Bond.")
	} else {
		fmt.Println("Neither")
	}

	switch x {
	case "Miss Moneypenny":
		fmt.Println("Hellooo James")
	case "James Bond":
		fmt.Println("Bond. James Bond.")
	default:
		fmt.Println("Neither")
	}

	// switch without expr:
	switch {
	case 10 > 20:
		fmt.Println("I do not know math")
	case 10 < 20:
		fmt.Println("I remember everything")
	default:
		fmt.Println("Why did i print?")
	}

}

/*
output:

Bond. James Bond.
Bond. James Bond.
I remember everything
*/
