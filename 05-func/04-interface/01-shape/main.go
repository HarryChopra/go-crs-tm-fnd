package main

import (
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() float64
}

type Square struct {
	Length float64
}

func (s Square) CalculateArea() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.Length * r.Width
}

func main() {
	s := Square{
		Length: 11.1,
	}
	printArea(s)

	c := Circle{
		Radius: 12.2,
	}
	printArea(c)

	r := Rectangle{
		Length: 10,
		Width:  5.5,
	}
	printArea(r)
}

func printArea(s Shape) {
	switch s.(type) {
	case Square:
		fmt.Printf("The area of the square is:\t%.2f\n", s.CalculateArea())
	case Circle:
		fmt.Printf("The area of the Circle is:\t%.2f\n", s.CalculateArea())
	case Rectangle:
		fmt.Printf("The area of the Rectangle is:\t%.2f\n", s.CalculateArea())
	default:
		fmt.Println("Unrecognised shape")
	}
}

/*
The area of the square is:      123.21
The area of the Circle is:      467.59
The area of the Rectangle is:   55.00
*/
