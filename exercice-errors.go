package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(err))
}

func SqrtE(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	iterations := 1
	for math.Abs((z*z-x)/x) > Precision {
		z -= (z*z - x) / (2 * z)
		iterations++
	}
	return z, nil
}

func exerciceErrors() {
	fmt.Println(SqrtE(2))
	fmt.Println(SqrtE(-2))
}
