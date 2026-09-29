package main

import (
	"fmt"
)

func main() {
	// /When the right hand side of the declaration is typed, the new variable is of that same type:
	var i int
	j := i // j is an int

	fmt.Println("Type of j is %T", j)	

	// But when the right hand side contains an untyped numeric constant, the new variable 
	// may be an int, float64, or complex128 depending on the precision of the constant:

	x := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("Type of x is %T\n", x)
	fmt.Printf("Type of f is %T\n", f)
	fmt.Printf("Type of g is %T\n", g)

}