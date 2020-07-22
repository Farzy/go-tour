package main

import "fmt"

func pointers() {
	i, j := 42, 2701

	fmt.Println("i =", i)
	p := &i
	fmt.Println("*p = ", *p)
	*p = 21
	fmt.Println("i =", i)

	p = &j
	fmt.Println("j = ", j)
	*p = *p / 37
	fmt.Println("j = ", j)
}

type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println("Struct (Println):", Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Printf("Struct (Printf %%V %%T) : %v of type %[1]T\n", v)
	fmt.Printf("Struct (Printf %%+V) : %+v\n", v)
	fmt.Printf("Struct (Printf %%#V) : %#v\n", v)

	v.X = 4
	fmt.Printf("new V = %v, v.X = %v\n", v, v.X)

	var p *Vertex
	p = &v

	p.X = 1e9
	fmt.Printf("new V via pointer = %v\n", v)

	// Struct literals
	v1 := Vertex{1, 2}
	v2 := Vertex{
		X: 1,
	}
	v3 := Vertex{}
	p_ := &Vertex{
		X: 1,
		Y: 2,
	}
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v, p = %v\n", v1, v2, v3, p_)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("#a = %#v\n", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("primes = %v\n", primes)
}

func moreTypes() {
	subtitle("Pointers")
	pointers()
	subtitle("Struct")
	structs()
	subtitle("Arrays")
	arrays()
}
