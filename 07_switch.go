package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("Write ", i, " as ")

	// Here’s a basic switch.
	switch i {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	default:
		fmt.Println("UNKNOWN")
	}

	// You can use commas to separate multiple expressions in the same case statement.
	// We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is weekday")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon", t.Hour())
	default:
		fmt.Printf("It's after noon %d:%d\n", t.Hour(), t.Minute())
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool %p\n", &i)
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}

	whatAmI(1)
	whatAmI(true)
	whatAmI("test")
	var arr [1_000_000]int
	whatAmI(arr)
}
