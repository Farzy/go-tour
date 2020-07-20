package main

import (
	"fmt"
	"math"
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

	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(g)
	fmt.Printf("Type conversion: int(%v, %v), float64(%v), uint(%v)\n", x, y, g, z)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.1
}

func constants() {
	const Pi float64 = 3.14
	const Pi2 = 3.14
	fmt.Printf("Pi = %v or %v\n", Pi, Pi2)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	fmt.Println("needInt(Small) =", needInt(Small))
	fmt.Println("needFloat(Small) =", needFloat(Small))
	fmt.Println("needFloat(Big) =", needFloat(Big))
	// Error: constant 1267650600228229401496703205376 overflows int
	//fmt.Println("needInt(Big) =", needInt(Big))
}
