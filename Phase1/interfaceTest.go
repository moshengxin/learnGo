package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
	test() float64
}

type Rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (r Rect) test() float64 {
	return 1
}
func (c circle) test() float64 {
	return 1
}

func main() {
	r := Rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(&r)
	measure(&c)
}
