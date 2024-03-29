package main

import "fmt"

// Branching with if and else in Go is straight-forward.
func main() {
	eight := 8

	// Here’s a basic example.
	if eight%2 == 0 {
		fmt.Printf("%d is even\n", eight)
	} else {
		fmt.Printf("%d is odd\n", eight)
	}

	// You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables declared in this statement are available in all branches.
	if n := 9; n < 0 {
		fmt.Println("n is negative")
	} else if n < 10 {
		fmt.Println("n has 1 digit")
	} else {
		fmt.Println("n has multiple digits")
	}

	fmt.Println("There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.")
}

// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
