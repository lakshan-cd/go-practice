package main

import (
	"fmt"
)

// /A defer statement defers the execution of a function until the surrounding function returns.

func deferPrint() {
	defer fmt.Println("This is sentence used with defer")
	fmt.Println("This is sentence used without defer")
}
