package main

import "fmt"

//note : type is after the variable name , then retun type

func add(x int, y int) int {
	return x + y
}


//if you have multiple parameters of the same type you can omit the type name for the first n-1 parameters
func addWithOnlyOneType(x, y int , z string) int {
	fmt.Println(z)
	return x + y
}

//
func retunAnyNuResult(x,y string) (string,string) {
	return x,y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addWithOnlyOneType(10, 13, "this is a string"))
	fmt.Println(retunAnyNuResult("this is first string1", "this is second string2"))
}
