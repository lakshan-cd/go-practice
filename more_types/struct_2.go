package main

import "fmt"

// A struct is a collection of fields.

func accessStructValues() {
	product := ProductPrice{22.25, 5.25}
	product.addedProfict = 10.25
	fmt.Print(product.addedProfict)
}
