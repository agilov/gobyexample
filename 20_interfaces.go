package main

import (
	"fmt"
	"math"
)

// *Interfaces* are named collections of method signatures
// read more: https://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go

// Here's a basic interface for geometric chapes
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implemet this interface on `rect` and `circle` types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type square struct {
	width float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
// Here we implement geometry on rects.
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (s *square) area() float64 {
	return math.Pow(s.width, 2)
}

func (s *square) perim() float64 {
	return 4 * s.width
}

// If a variable has an interface type, then we can call all methods
// that are in the named interface.
// Here's a generic `measure` function taking advantage of this to work on any `geometry`
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func measurePointer(gp geometry) {
	fmt.Println(gp)
	fmt.Println("measurePointer area", gp.area())
	fmt.Println("measurePointer perim", gp.perim())
}

func main() {
	c := circle{radius: 10}
	s := square{width: 5}
	r := rect{width: 12, height: 12}

	// The circle and rect struct types both implement the geometry interface
	// so we can use instances of these structs as arguments to measure.
	measure(c)
	measurePointer(&s)
	measure(r)
}

// To learn more about Go’s interfaces,
// check out this (great blog post)[https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go].
