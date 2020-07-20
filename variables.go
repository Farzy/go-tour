package main

import "fmt"

var c, python, java bool
var j, k int = 1, 2

func variables() {
	var i int
	fmt.Println(i, j, k, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(c2, python2, java2)
}
