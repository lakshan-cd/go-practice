package main

import "fmt"

// /This function returns an integer, and I'll call that returned value result.

func add(x , y int)(xVariable int,result int){
	xVariable = x
	result = x + y
	return
}

func main() {
	fmt.Println(add(42, 13))
}