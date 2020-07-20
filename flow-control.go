package main

import (
	"fmt"
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
