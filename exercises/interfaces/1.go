package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}

type Rectangle struct {
	l float64
	b float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.b
}

func PrintArea(s Shape) {
	fmt.Println("Area of the shape:", s.Area())
}

func main() {
	c := Circle{
		r: 10,
	}

	r := Rectangle {
		l: 5,
		b: 10,
	}

	PrintArea(c)
	PrintArea(r)
}
