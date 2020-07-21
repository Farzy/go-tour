package main

import "fmt"

func defer1() {
	defer fmt.Println("…World")

	fmt.Println("Hello…")
}

func defer2() {
	fmt.Println("Couting…")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
func defers() {
	defer1()
	defer2()
}
