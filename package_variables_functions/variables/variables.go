package variables

import "fmt"

//varibales can be declared using var, can be package or function level

var c, python , java bool //all these are boolean

func main() {
	var i int //variables are automatically given a zero value when you declare them without assigning a value.
	fmt.Println(i, c, python, java)
}

// Variables declared without an explicit initial value are given their zero value.

// The zero value is:
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
