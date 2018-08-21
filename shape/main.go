package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	length float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {

	tri := triangle{base: 1, height: 2}
	square := square{length: 2, height: 4}

	printArea(tri)
	printArea(square)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.length * s.height
}

func printArea(s shape) {
	fmt.Println("The area is", s.getArea())
}
