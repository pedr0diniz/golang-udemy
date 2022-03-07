package main

import "fmt"

func main() {
	s := square{2.5}
	t := triangle{4, 5}
	printArea(s)
	printArea(t)
}

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
