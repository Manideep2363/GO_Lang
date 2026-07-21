package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("area:", s.Area())
}

func main() {
	PrintArea(Rectangle{Length: 10, Width: 5})

	PrintArea(Circle{Radius: 7})

}
