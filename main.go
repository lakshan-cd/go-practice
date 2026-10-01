// This is about packages
//and to run the starting point is main function
package main

import (
	"fmt"
	"math/rand"
)

func packagesIntro() {
	fmt.Println("My faviorite number is" , rand.Intn(10))
}

// Run everything in this folder with:
//   go run .
func main() {
	fmt.Println("--- main.go: packages ---")
	packagesIntro()

	fmt.Println("--- imports.go: imports ---")
	importsDemo()

	fmt.Println("--- exportedNames.go: exported names ---")
	exportedNamesDemo()
}
