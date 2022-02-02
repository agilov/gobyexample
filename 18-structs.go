package main

import "fmt"

// Go's `structs` are typed collections of fields.
// They're useful for grouping data together to form records.

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new `person` with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	return &p
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Roman", 30})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 20})

	// Ommited fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yelds a pointer to the struct
	fmt.Println(&person{name: "Beep", age: 11})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("John"))

	// Access struct fields with a dot.
	s := person{"Sean", 60}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.name)

	// But manual dereferencing is also possible.
	fmt.Println((*sp).age)

	// Structs are mutable
	sp.age = 61
	fmt.Println(sp.age)
}
