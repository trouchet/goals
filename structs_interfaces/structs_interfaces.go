package main

import (
	"fmt"
	"math"
)

// Shape is an interface for geometric shapes
type Shape interface {
	Area() float64
}

// Circle is a struct representing a circle
type Circle struct {
	Radius float64
}

// Rectangle is a struct representing a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// PrintArea prints the area of a shape
func PrintArea(s Shape) {
	fmt.Printf("Area of the shape: %f\n", s.Area())
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Use the PrintArea function with different shapes
	PrintArea(circle)
	PrintArea(rectangle)
}
