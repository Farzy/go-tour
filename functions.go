package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Group parameters of same type
func sub(x, y int) int {
	return x - y
}

// Return more than one value
func swap(x, y string) (string, string) {
	return y, x
}

func functions() {
	fmt.Printf("%d + %d = %d\n", 42, 13, add(42, 13))
	fmt.Printf("%d - %d = %d\n", 42, 13, sub(42, 13))
	a, b := swap("world", "Hello")
	fmt.Println("Swapped vars:", a, b)
}
