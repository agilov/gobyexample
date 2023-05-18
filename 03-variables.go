package main

import "fmt"

// In Go, variables are explicitly declared and used by
// the compiler to e.g. check type-correctness of function calls.
func main() {
	// var declares 1 or more variables.
	var a = "Merhaba!"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c = 5, 6
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	//For example, the zero value for an int is 0.
	var e uint
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	x := "banana"
	fmt.Println(x)
	
	// What happens if I change variable value?
	z := "apple"
	fmt.Printf("%s - %p \n", z, &z)
	
	z = "wood"
	fmt.Printf("%s - %p \n", z, &z)
}
