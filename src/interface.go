package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	height float64
	width  float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())	
}
