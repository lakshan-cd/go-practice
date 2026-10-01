package main

import "fmt"

//Short inizialize variables, inside only func, outside func we can use var keyword

// i and j live in variables2.go. Package level vars are shared by every file in the
// package, so declaring them again here would be a "redeclared" compile error.
// test := "this is a test" //short variable declaration, only inside func
// //retun -> .\variables3.go:8:1: syntax error: non-declaration statement outside function body

func shortDeclaration() {
	var c , python , java = true , false , "no!" //type is inferred from the value
    hello := "hello" //short variable declaration, only inside func
	fmt.Println(i, j, c, python, java, hello)
}