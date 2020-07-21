package main

import (
	"fmt"
	"math"
)

// Precision Stop iterating on Sqrt when relative difference is smaller than this
const Precision = 1e-14

// Sqrt calculates square root using Newton's method
// Stop iterating when relative difference of squares is less than Precision
func Sqrt(x float64) (float64, int) {
	z := 1.0
	iterations := 1
	for math.Abs((z*z-x)/x) > Precision {
		z -= (z*z - x) / (2 * z)
		iterations++
	}
	return z, iterations
}

// Sqrt2 calculates square root using Newton's method
// Stop iterating when relative difference of 2 consecutives results is less than Precision
func Sqrt2(x float64) (float64, int) {
	z, zPrev := 1.0, 0.0
	iterations := 0
	for (math.Abs(z-zPrev) / z) > Precision {
		zPrev = z
		z -= (z*z - x) / (2 * z)
		iterations++
	}
	return z, iterations
}

func exerciceLoopsAndFunctions() {
	res, iter := Sqrt(2)
	fmt.Printf("Sqrt(2) = %v (%v iterations)\n", res, iter)
	res, iter = Sqrt(2)
	fmt.Printf("Sqrt2(2) = %v (%v iterations)\n", res, iter)

	var vals = []float64{1, 4, 3, 9, 5, 25, 100, 1000}

	for _, x := range vals {
		res1, iter1 := Sqrt(x)
		res2, iter2 := Sqrt2(x)
		fmt.Printf("Sqrt(%v) = %v (%v iterations) / Sqrt2(%v) = %v (%v iterations) / math.Sqrt(%v) = %v\n",
			x, res1, iter1, x, res2, iter2, x, math.Sqrt(x))
	}
}
