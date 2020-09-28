package main

import (
	"fmt"
	"math"
)

type reat struct {
	width, height float64
}

func (r reat) area() float64 {
	return r.width * r.height
}

func (r reat) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := reat{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
