package main

import "fmt"

// range iterates over elements in a variety of data structures.
// Let’s see how to use range with some of the data structures we’ve already learned.
func main() {
	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn't need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though
	for i, v := range nums {
		if v == 3 {
			fmt.Println("index of 3:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate ofer just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune
	// and the second the rune itself
	for i, c := range " !\"#$%&'()*+,-./" {
		fmt.Println(i, c)
	}
}
