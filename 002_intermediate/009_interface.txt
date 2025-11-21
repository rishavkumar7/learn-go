package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

type Square struct {
	side float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) diameter() float64 {
	return 2 * c.radius
}

func printer(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func custom1(i interface{}) { // we can pass a single argument of any type to this function.
	switch i.(type) {
	case int:
		fmt.Println("This is an integer.")
	case string:
		fmt.Println("This is a string.")
	case float64:
		fmt.Println("This is a float64")
	default:
		fmt.Println("This is an unknown type.")
	}
}

func custom2(i ...interface{}) { // we can pass any number of arguments of any type to this function.
	for _, v := range i {
		fmt.Println(v)
	}
}

func main() {
	rectangle := Rectangle{
		length: 20.0,
		width:  12.0,
	}

	// square := Square{
	// 	side: 15.0,
	// }

	circle := Circle{
		radius: 10.0,
	}

	printer(&rectangle) // Rectangle struct implements the Geometry interface.
	// printer(&square) // Square struct does not implement the Geometry interface as the perimeter method is missing in Square struct.
	printer(circle) // Circle struct implements the Geometry interface.

	custom1("Rishav")
	custom1(26)
	custom1(true)

	custom2("Rishav", true, 62.8, 26)
}

// Interface is used for decoupling, polymorphism, etc.
// Passing the rectangle as a reference because the perimeter() method of Rectangle struct has a pointer receiver and this method is required to implement the Geometry interface.
// Passing the circle as a value because neither the area() method nor the perimeter() method of Circle struct has a pointer receiver. Only diameter() method has a pointer receiver but its irrelevant in terms of implementing the Geometry interface.
