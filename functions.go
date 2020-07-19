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

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// Naked return
	return
}

func functions() {
	fmt.Printf("%d + %d = %d\n", 42, 13, add(42, 13))
	fmt.Printf("%d - %d = %d\n", 42, 13, sub(42, 13))
	a, b := swap("world", "Hello")
	fmt.Println("Swapped vars:", a, b)
	x, y := split(42)
	fmt.Println("Named return values, split(42) = ", x, y)
}
