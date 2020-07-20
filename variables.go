package main

import (
	"fmt"
	"math/cmplx"
)

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

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//noinspection GoBoolExpressions
func types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Default values: %v %v %v %q\n", i, f, b, s)
}
