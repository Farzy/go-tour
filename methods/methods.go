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

func Main() {
	v := Vertex{3, 4}
	fmt.Printf("Abs(%v) = %v\n", v, v.Abs())
}
