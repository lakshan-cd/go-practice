package main

import (
	"fmt"
	"math"
)

// named powShort because if_3.go already has a pow in this package
func powShort(x , n, lim float64) float64 {
	if v := math.Pow(x , n);v < lim {
		return v
	}
	return lim

}

func ifShortStatement() {
	fmt.Println(
		powShort(3, 2, 10),
		powShort(3, 3, 20),
	)
}