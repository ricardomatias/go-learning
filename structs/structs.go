package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	c := Circle{0, 0, 5}
	roger := Person{"Roger"}

	a := Android{roger, "LOOKS"}

	a.Talk()

	fmt.Println("area", c.area())
}
