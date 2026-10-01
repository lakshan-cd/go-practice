package main

import "fmt"

// A struct literal denotes a newly allocated struct value by listing the values of its fields.

type Test struct {
	X, Y int
}

var (
	v1 = Test{1, 2}  // has type Test
	v2 = Test{X: 1}  // Y:0 is implicit
	v3 = Test{}      // X:0 and Y:0
	p  = &Test{1, 2} // has type *Test
)

func structLiterals() {
	fmt.Println(v1, v2, v3, p)
}
