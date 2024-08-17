package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	s := square{sideLength: 10}
	r := triangle{base: 10, height: 10}

	printArea(s)
	printArea(r)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
