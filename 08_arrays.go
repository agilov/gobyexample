package main

import "fmt"

func main() {
	// Here we create an array a that will hold exactly 5 ints.
	// The type of elements and length are both part of the array’s type.
	// By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("empty array of 5 ints:", a)

	// We can set a value at an index using the array[index] = value syntax,
	// and get a value with array[index].
	a[3] = 10000000
	fmt.Println("set value of 4'th element:", a)
	fmt.Println("value of 4'th element:", a[3])

	// The builtin len returns the length of an array.
	len := len(a)
	fmt.Printf("array length: %T %d\n", len, len)

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, -1, -2}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	const width int = 5
	const height int = 5
	const length int = 5

	var twoD [width][height][length]int
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			for k := 0; k < length; k++ {
				twoD[i][j][k] = k + i + j
			}
		}
	}

	// Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
	fmt.Println("twoD:", twoD)
}
