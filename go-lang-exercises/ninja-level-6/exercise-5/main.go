/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
}

type Square struct {
	l int
	w int
}

func (s Square) area() int {
	return s.l * s.w
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func main() {
	square := Square{
		20,
		30,
	}

	circle := Circle{
		30,
	}

	fmt.Println(circle.area())
	fmt.Println(square.area())
}
