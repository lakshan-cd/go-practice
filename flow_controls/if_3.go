package main

import (
	"fmt"
	"math"
)

//Variables declared inside an if short statement are also available inside any of the else blocks.



func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) //%g -> print a floating-point number in a compact form.
	}
	// can't use v here, though
	return lim
}

func ifElse() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}