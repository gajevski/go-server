package main

import (
	"fmt"
	"math"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func CircleFactory(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func (circle Circle) calculateCircumference() {
	circumference := 2 * PI * circle.radius
	fmt.Println(circumference)
}

func (circle Circle) calculateArea() {
	area := PI * math.Pow(circle.radius, 2)
	fmt.Println(area)
}

func main() {
	circle := CircleFactory(1.5)
	circle.calculateCircumference()
	circle.calculateArea()
}
