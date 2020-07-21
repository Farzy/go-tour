package main

import "fmt"

func func1() int {
	defer fmt.Println(
		"…World (this is printed after func2's counter: the return value is " +
			"calculated but held until func1's defer functions are executed.)")

	fmt.Println("func1: Hello…")

	return func2()
}

func func2() int {
	fmt.Println("func2: Counting…")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	return 42
}
func defers() {
	func1()
	//func2()
}
