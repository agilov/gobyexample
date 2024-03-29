package main

import "fmt"

func main() {
	i := 1

	// The most basic type, with a single condition.
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 4; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly
	// until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 10; n >= 0; n-- {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
