package main

import (
	"fmt"
	"math"
	"math/rand"
)

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum(0..9) =", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum(x2 until 1000) =", sum)

	for {
		r := rand.Intn(1000)
		fmt.Println("Loop forever… almost:", r)
		if r < 42 {
			break
		}
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func ifs() {
	fmt.Println("pow(3, 2, 10) =", pow(3, 2, 10))
	fmt.Println("pow(3, 3, 20) =", pow(3, 3, 20))
}
