package main

import "fmt"

// Run everything in this folder with:
//   go run ./package_variables_functions/variables
func main() {
	fmt.Println("--- variables.go: zero values ---")
	zeroValues()

	fmt.Println("--- variables2.go: omit the type when you initialize ---")
	inferredTypes()

	fmt.Println("--- variables3.go: short declaration ---")
	shortDeclaration()
}
