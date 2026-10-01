package main

import "fmt"

// A struct is a collection of fields.

type ProductPrice struct {
	unitPrice    float64
	addedProfict float64
}

func printStruct() {
	fmt.Println(ProductPrice{22.25, 5.25})
}
