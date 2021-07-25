package main

import "fmt"

type Vehicle struct {
	Doors int
	Color string
}

type Truck struct {
	Vehicle
	FourWheel bool
}

type Saloon struct {
	Vehicle
	Luxury bool
}

func main() {
	v1 := Truck{
		Vehicle: Vehicle{
			Doors: 6,
			Color: "Black",
		},
		FourWheel: true,
	}
	fmt.Printf("v1: %#v\n", v1)
	fmt.Println("v1's color:", v1.Color)

	v2 := Saloon{
		Vehicle: Vehicle{
			Doors: 4,
			Color: "Red",
		},
		Luxury: false,
	}
	fmt.Printf("v2: %#v\n", v2)
	fmt.Println("v2's luxury:", v2.Luxury)
}

/*
v1: main.Truck{Vehicle:main.Vehicle{Doors:6, Color:"Black"}, FourWheel:true}
v1's color: Black

v2: main.Saloon{Vehicle:main.Vehicle{Doors:4, Color:"Red"}, Luxury:false}
v2's luxury: false
*/
