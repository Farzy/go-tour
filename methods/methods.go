package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method defined on type local to the package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt(2))
	fmt.Printf("Method Abs(%v) = %v\n", v, v.Abs())
	fmt.Printf("Function Abs(%v) = %v\n", v, Abs(v))
	fmt.Printf("MyFloat.Abs(%v) = %v\n", f, f.Abs())
}
