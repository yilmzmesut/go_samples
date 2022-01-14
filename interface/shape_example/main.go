package main

import "fmt"

type square struct {
	sideLengt float64
}
type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	square := square{sideLengt: 10}
	triangle := triangle{base: 10, height: 10}
	printArea(square)
	printArea(triangle)
}

func (s square) getArea() float64 {
	return s.sideLengt * s.sideLengt
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
