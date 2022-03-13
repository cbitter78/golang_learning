package main

import (
	"fmt"
	"math"
)

type squar struct {
	side float64
}

type rectangle struct {
	lenight float64
	width   float64
}

func (r rectangle) area() float64 {
	return r.lenight * r.width
}

func (s squar) area() float64 {
	return s.side * 2
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("The Area is:", s.area())
}

func main() {
	fmt.Println("Ninja level 6 Exercises 5:")
	c := circle{
		radius: 12.345,
	}
	s := squar{
		side: 18.67,
	}
	r := rectangle{
		lenight: 20,
		width:   30,
	}
	info(c)
	info(s)
	info(r)
}
