package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func functions() {
	fmt.Printf("%d + %d = %d\n", 42, 13, add(42, 13))
}
