package main

import "fmt"

// Run everything in this folder with:
//   go run ./package_variables_functions/types
func main() {
	fmt.Println("--- basicTypes.go: Go's basic types ---")
	basicTypes()

	fmt.Println("--- typeInference.go: type inference ---")
	typeInference()

	fmt.Println("--- typeConversion.go: type conversion ---")
	typeConversion()

	fmt.Println("--- constants.go: constants ---")
	constants()
}
