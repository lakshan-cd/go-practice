package main

import (
	"fmt"
)

func constants() {
	// can be string, boolean, or numeric values.
	const Pi = 3.14 //cannot be assing using := operator, must be assigned using = operator
	fmt.Println("Value of Pi is", Pi)
}