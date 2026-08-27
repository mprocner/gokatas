// Package sqrt calculates square root: given number x, we want to find the
// number z for which z² is most nearly x. Source: go.dev/tour/flowcontrol/8
//
// Level: beginner
// Topics: algorithms, math
package main

import (
	"fmt"
	"math"
)

const delta = 1e-15

func main() {
	fmt.Println(sqrt(3))
}

func sqrt(x float64) float64 {
	z := 1.0 // initial guess
	for math.Abs(z*z-x) > delta {
		z -= (z*z - x) / (2 * z) // Newton's method
	}
	return z
}
