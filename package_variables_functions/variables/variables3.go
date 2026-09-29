package variables

import "fmt"

//Short inizialize variables, inside only func, outside func we can use var keyword

var i, j int = 10 , 20 
// test := "this is a test" //short variable declaration, only inside func
// //retun -> .\variables3.go:8:1: syntax error: non-declaration statement outside function body

func main() {
	var c , python , java = true , false , "no!" //type is inferred from the value
    hello := "hello" //short variable declaration, only inside func
	fmt.Println(i, j, c, python, java, hello)
}