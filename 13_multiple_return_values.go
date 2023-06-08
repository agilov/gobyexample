package main

import "fmt"

// Go has built-in support for *multiple return values*.
// This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

// The `(int, int)` in this function signature shows
// that the function returns 2 `ints`.
func vals() (int, int, string) {
	b := 1
	a := "test"
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

	return b, 2, a
}

func main() {
	// Here we use the 2 different return values from the call
	// with the *multiple assignment*.
	a, b, y := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(y)

	// If you only want a subset of the returned values,
	// use the blank identifier _.
	_, c, k := vals()
	fmt.Println(c)
	fmt.Printf("%p\n", &k)
}

// Accepting a variable number of arguments is another nice feature of Go functions;
// we'll look at this next
