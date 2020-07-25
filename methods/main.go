package methods

import (
	"fmt"
	"go-tour/utils"
	"math"
)

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
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

	// Interface
	utils.Subtitle("Interface")
	var a Abser

	a = f  // MyFloat implements Abser
	a = &v // *Vertex implements Abser
	//a = v	// Vertex does not implement Abser

	fmt.Printf("a is an interface var (= &v): %v\n", a.Abs())

	var i I = T{"Hello"}
	i.M()

	j := &T{"Hello"}
	describe(j)
	j.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	utils.Subtitle("Interface values with nil underlying values")
	nils()

	utils.Subtitle("Nil interface values")
	nilInterfaceValue()
}
