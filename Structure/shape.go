package main

import "math"

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func Perimeter(r Rectangle) float64 {

	return 2 * (r.length + r.breadth)
}

func (r Rectangle) Area() float64 {
	return r.breadth * r.length

}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
