package main

import "fmt"

type rect struct {
	width, height int
}

func newRect(width, height int) *rect {
	rect := rect{width, height}

	return &rect
}

func(r *rect) area() int {
	return r.width * r.height
}

func(r *rect) perim() int {
	return (2 * r.width) + (2 * r.height)
}

func main() {
	rectangle := newRect(3, 5)

	fmt.Println("initialized rect", rectangle)
	fmt.Println("area:", rectangle.area())
	fmt.Println("perimeter:", rectangle.perim())
}
