package main

import "fmt"

// Go supports [pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming)),
// allowing you to pass references to values and records within your programm

// We'll show how pointers work in contrast to values with 2 functions: `zeroval` and `zeroptr`.
// `zeroval` has an `int` parameter so arguments will be passed to it by value.
// `zeroval` will get a copy of `ival` distinct from the on in the calling function.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning that it takes an `int` pointer.
// The `*iptr` code in the function body then *dereferences* the pointer from it's
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the refrerenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("init:", i)

	zeroval(i)
	fmt.Println("zval:", i)

	// The `&i` syntax gives the memory address of `i`, i.e. pointer to `i`.
	zeroptr(&i)
	fmt.Println("zptr:", i)

	// Pointers can be printed too
	fmt.Println("ptr:", &i)

	// `zeroval` doesn't change the `i` in `main`, but `zeroptr` does because it
	// has memory address for that variable
}
