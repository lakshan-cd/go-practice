package main

import (
	"fmt"
	"time"
)

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
func switchCase() {
	fmt.Println("When's Staurday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In tow days.")
	default:
		fmt.Println("Too far away.")
	}
}
