package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Main() {
	v := Vertex{3, 4}
	fmt.Printf("Method Abs(%v) = %v\n", v, v.Abs())
	fmt.Printf("Function Abs(%v) = %v\n", v, Abs(v))
}
