package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) AbsV() float64 {
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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt(2))
	fmt.Printf("Method Abs(%v) = %v\n", v, v.Abs())
	fmt.Printf("Function Abs(%v) = %v\n", v, Abs(v))
	fmt.Printf("MyFloat.Abs(%v) = %v\n", f, f.Abs())
	v.Scale(10)
	fmt.Printf("Scale(v) = %v\n", v)
	Scale(&v, 10)
	fmt.Printf("Scale(v) as function = %v\n", v)
	p := &v
	p.Scale(2)
	fmt.Printf("Scale(&p) = %v\n", v)
	Scale(p, 10)
	fmt.Printf("Scale(&p) as function = %v\n", v)
	// Functions do not always accept pointers instead of values
	//fmt.Printf("Function Abs(p) = %v\n", Abs(&v))

	w := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", w, w.Abs())
	w.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", w, w.Abs())
}
