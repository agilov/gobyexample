package main

import "fmt"

// Go supports *methods* defined on struct types

type rect struct {
	width, height int
}

// This `area` method has a *receiver* type of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// This method wont mutate actual struct because struct will be copied to receive this method
// Instead it returns the copy with mutated value.
func (r rect) SetWidth(w int) rect {
	r.width = w
	return r
}

// This method will mutate it's receiver in the outer scope because receiver is pointer to struct
func (r *rect) MutateWidth(w int) {
	r.width = w
}

type me struct {
	weight int
}

func (me *me) getFat(kg int) {
	me.weight += kg
}
func (me me) getMoreFat(kg int) me {
	me.weight += kg
	return me
}

func main() {

	r := rect{width: 10, height: 11}
	fmt.Println("rect:", r)

	// Here we call 2 methods defined for our struct.
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls
	// or to allow the method to mutate the receiver struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	// SetWidth won't mutate `r` struct because it is copied for the call.
	fmt.Println("before set:", r)
	nr := r.SetWidth(5)
	fmt.Println("old:", r)
	fmt.Println("new:", nr)

	// MutateWidth will mutate `r` because it receives call as a reference
	fmt.Println("before mutate:", r)
	r.MutateWidth(5)
	fmt.Println("mutated:", r)

	me := me{weight: 10}
	me.getFat(20)
	k := me.getMoreFat(20)
	fmt.Println("me:", me)
	fmt.Println("k:", k)
}

// Next we'll look at Go's mechanism for grouping and
// naming related sets of methods: interfaces
