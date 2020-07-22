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

// Vertex is a sample struct
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
	pp := &Vertex{
		X: 1,
		Y: 2,
	}
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v, p = %v\n", v1, v2, v3, pp)
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

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d s=%v\n", len(s), cap(s), s)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Printf("Slice [1:4] of prime = %v\n", s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names =", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Printf("a = %v, b = %v\n", a, b)

	b[0] = "XXX"
	fmt.Println("b[0] = \"XXX\"")
	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Println("names =", names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("q = %v (%#[1]v)\n", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Printf("r = %v (%#[1]v)\n", r)

	u := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("s = %v\n", u)

	q = q[1:4]
	fmt.Printf("q[1:4] = %v\n", q)

	q = q[:2]
	fmt.Printf("q[:2] = %v\n", q)

	q = q[1:]
	fmt.Printf("q[1:] = %v\n", q)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	println("Slice the slice to give it zero length")
	s = s[:0]
	printSlice(s)

	fmt.Println("Extend its length")
	s = s[:4]
	printSlice(s)

	fmt.Println("Drop its first 2 values")
	s = s[2:]
	printSlice(s)

	fmt.Println("Extend length to 3")
	s = s[:3]
	printSlice(s)

	// panic: runtime error: slice bounds out of range [:5] with capacity 4
	//fmt.Println("Extend length to 5")
	//s = s[:5]
	//printSlice(s)

	// Nil slice
	var ns []int
	fmt.Println("Nil slice")
	printSlice(ns)
	//noinspection GoNilness
	if ns == nil {
		fmt.Println("s is nil")
	}
}

func moreTypes() {
	subtitle("Pointers")
	pointers()
	subtitle("Struct")
	structs()
	subtitle("Arrays")
	arrays()
	subtitle("Slices")
	slices()
}
