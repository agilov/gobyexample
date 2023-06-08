package main

import "fmt"

// Go has various value types including strings, integers, floats, booleans, etc.
// Here are a few basic examples.
func main() {

	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("79+66=", 79+66)
	fmt.Println("55.5/12.4=", 55.5/12.4)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
