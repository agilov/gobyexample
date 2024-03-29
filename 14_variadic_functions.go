package main

import "fmt"

// [Variadic functions](https://en.wikipedia.org/wiki/Variadic_function) can be
// called with any number of trailing arguments.
// For example `fmt.Prinln` is a common variadic function.

// Here's a function that will take an arbitrary number of `ints` as arguments.
func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)

	return total
}

func main() {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	//nums := []int{1, 2, 3, 4, 5}
	sum([]int{123, 123, 1, 3, 4}...)
}

// Another key aspect of functions in Go is their ability to form closures,
// which we’ll look at next.
