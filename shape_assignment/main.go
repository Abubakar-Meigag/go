package main

import "fmt"


type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}

type triangle struct{
	height float64
	base float64
}

func main() {
	s := square{sideLength: 5}
	t := triangle{height: 3, base: 4}

	printArea(s)
	printArea(t)
	fmt.Println("Square area:", s.getArea())
	fmt.Println("Triangle area:", t.getArea())
	fmt.Println("Total area:", s.getArea() + t.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
