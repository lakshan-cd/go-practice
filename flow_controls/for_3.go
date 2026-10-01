package main

import "fmt"

//in go for is the while loop , we can drop the semicolns
func forAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}