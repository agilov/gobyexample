package main

import (
	"fmt"
	"runtime"
	"time"
)

// Go supports [anonymous functions](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form [closures](https://en.wikipedia.org/wiki/Closure_(computer_programming))
// Anonymous functions are useful when you want to define a function inline
// without having to name it.

var world *int

// This function `intSeq` return another function, which we define anonymously
// in the body of `intSeq`. The returned function *closes over* the variable `i`
// to form a closure.
func intSeq() func() func() func() func() func() func() int {
	i := 0
	return func() func() func() func() func() func() int {
		j := 0
		return func() func() func() func() func() int {
			k := 0
			return func() func() func() func() int {
				z := 0
				return func() func() func() int {
					y := 0
					return func() func() int {
						w := 0
						return func() int {
							w++
							y++
							z++
							k++
							j++
							i++
							fmt.Printf("%p %p %p %p %p %p\n", &i, &j, &k, &z, &y, &w)
							world = &i
							return i + j + k + z + y + w
						}
					}
				}
			}
		}
	}
}

func subMain() {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value,
	// which will be updated each time we call nextInt.
	nextInt := intSeq()()()()()()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Printf("%p\n", world)

	// To confirm that the state is unique to that particular function,
	// create and test a new one.
	newSeq := intSeq()
	fmt.Println(newSeq()()()()()())
	fmt.Printf("%p\n", world)
}

func main() {
	subMain()
	runtime.GC()
	time.Sleep(time.Second)
	*world = 131313123123
	fmt.Printf("%p\n", world)
	fmt.Printf("Closure still exist: %d\n", *world)
}

// The last feature of functions we'll look at for now is recursion
