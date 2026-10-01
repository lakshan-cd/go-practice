package main

import "fmt"

// Go has only one looping construct, the for loop.
func forBasic() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
 }

//  the 3 parts are
// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration