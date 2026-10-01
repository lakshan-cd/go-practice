package main

import (
	"fmt"
)

// Run everything in this folder with:
//
//	go run ./flow_controls
func main() {
	fmt.Println("--- for_1.go: the basic for loop ---")
	forBasic()

	fmt.Println("--- for_2.go: init and post are optional ---")
	forOptionalInitPost()

	fmt.Println("--- for_3.go: for as a while loop ---")
	forAsWhile()

	// for_4.go: forForever() loops forever, uncomment only if you want to Ctrl+C out
	// forForever()

	fmt.Println("--- if_1.go: if without parentheses ---")
	ifBasic()

	fmt.Println("--- if_2.go: if with a short statement ---")
	ifShortStatement()

	fmt.Println("--- if_3.go: if / else ---")
	ifElse()

	fmt.Println("--- switch_1.go: switch ---")
	switchBasic()

	fmt.Println("--- switch_2.go: switch evaluation order ---")
	switchCase()

	fmt.Println("--- switch_3.go: switch without a condition ---")
	switchWithoutCondition()

	fmt.Println("--- defer_1.go: defer ---")
	deferPrint()

	fmt.Println("--- defer_2.go: defer stack ---")
	deferStack()
}
