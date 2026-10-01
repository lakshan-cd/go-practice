package main

import "fmt"

// /This function returns an integer, and I'll call that returned value result.

// named addNamed because functions.go already has an add in this package
func addNamed(x , y int)(xVariable int,result int){
	xVariable = x
	result = x + y
	return
}

func namedReturnValues() {
	fmt.Println(addNamed(42, 13))
}