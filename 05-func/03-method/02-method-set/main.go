package main

import (
	"fmt"
	"math"
)

type Shape interface {
	calculateArea() float64
}

type Square struct {
	Length float64
}

func (s *Square) calculateArea() float64 {
	return s.Length * s.Length
}

type Rectangle struct {
	Length, Width float64
}

func (r *Rectangle) calculateArea() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c *Circle) calculateArea() float64 {
	return 2 * math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	switch s.(type) {
	case *Square:
		fmt.Printf("The area of square is:\t\t%.2f\n", s.calculateArea())
	case *Rectangle:
		fmt.Printf("The area of rectangle is:\t%.2f\n", s.calculateArea())
	case *Circle:
		fmt.Printf("The area of circle is:\t\t%.2f\n", s.calculateArea())
	}
}

func main() {
	s := &Square{5}
	r := &Rectangle{10, 11}
	c := &Circle{3.9}

	printArea(s)
	printArea(r)
	printArea(c)
}
