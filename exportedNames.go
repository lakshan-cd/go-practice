//if want to export that should be statrt with capital letter

package main

import (
	"fmt"
	"math"
)

// func main() {
// 	fmt.Println(math.pi)
// }

///the abouve one throes this

// C:\Users\LakshanChamara\Desktop\New folder\lakshan-cd\go lang\practice>go run exportedNames.go
// # command-line-arguments
// .\exportedNames.go:11:19: undefined: math.pi (but have Pi)

// C:\Users\LakshanChamara\Desktop\New folder\lakshan-cd\go lang\practice>


//but this work

func exportedNamesDemo() {
	fmt.Println(math.Pi)
}