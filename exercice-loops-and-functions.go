package main

import (
	"fmt"
	"math"
)

const Precision = 1e-14

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs((z*z-x)/x) > Precision {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func exerciceLoopsAndFunctions() {
	fmt.Printf("Sqrt(2) = %v\n", Sqrt(2))

	var vals = []float64{1, 4, 3, 9, 5, 25}

	for _, x := range vals {
		fmt.Printf("Sqrt(%v) = %v / math.Sqrt(%v) = %v\n", x, Sqrt(x), x, math.Sqrt(x))
	}
}
