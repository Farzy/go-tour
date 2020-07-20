package main

import "fmt"

var c, python, java bool
var j, k int = 1, 2

//noinspection GoBoolExpressions
func variables() {
	var i int
	fmt.Println(i, j, k, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(c2, python2, java2)

	l := 3
	c3, python3, java3 := true, false, "no!"
	fmt.Println(l, c3, python3, java3)
}
