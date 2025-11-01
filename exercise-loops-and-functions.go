package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	eps := 0.00001
	for i := 0; math.Abs(x-z*z) > eps; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d Value %f\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
