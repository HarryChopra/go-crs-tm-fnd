package main

import "fmt"

func main() {
	p1 := struct {
		First, Last     string
		FavoriteFlavors []string
		MoviesReleases  map[string]int
	}{
		First:           "Ian",
		Last:            "Fleming",
		FavoriteFlavors: []string{"Chocolate", "Martini", "Vanilla"},
		MoviesReleases: map[string]int{
			"Dr. No":                1962,
			"From Russia With Love": 1963,
			"GoldFinger":            1964,
			"Thunderball":           1965,
			"You only live twice":   1967,
		},
	}
	fmt.Println("Name:", p1.First, p1.Last)
	for _, fl := range p1.FavoriteFlavors {
		fmt.Printf("Favorite Flavors: %s,", fl)
	}
	fmt.Println("\nMovieReleases:")
	for k, v := range p1.MoviesReleases {
		fmt.Printf("\t%d: %s\n", v, k)
	}
}

/*
Name: Ian Fleming
Favorite Flavors: Chocolate,Favorite Flavors: Martini,Favorite Flavors: Vanilla,
MovieReleases:
        1962: Dr. No
        1963: From Russia With Love
        1964: GoldFinger
        1965: Thunderball
        1967: You only live twice
*/
