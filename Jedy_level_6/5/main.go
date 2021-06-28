package main

import (
	"fmt"
	"math"
)

func main() {
	c := circle{
		radius: 5,
	}

	s := square{
		side: 8,
	}

	info(c)
	info(s)

}

type shape interface {
	area() float64
}

type circle struct {
	radius int
}

type square struct {
	side int
}

func (f circle) area() float64 {
	return math.Pi * math.Pow(float64(f.radius), 2)
}

func (f square) area() float64 {
	return math.Pow(float64(f.side), 2)
}

func info(s shape) {
	fmt.Println(s.area())
}
