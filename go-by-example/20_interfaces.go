package main

import (
	"fmt"
	"math"
)

// interfaces are like inheritance

type geometry interface {
	area() float64	// these are functions that will be overloaded later
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func(r rect) area() float64 {	// member function takes in rect and returns area
	return r.width * r.height
}
func(r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

func(c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func(c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {	// we can use these overloaded structs like this generally (abstraction)
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	rect := rect{12.1, 6.4}
	circle := circle{2.4}

	fmt.Println("measuring rectangle")
	measure(rect)

	fmt.Println("measuring circle")
	measure(circle)

}
