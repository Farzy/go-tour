package main

import (
	"fmt"
	"strings"
)

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

func printSlice(s string, x []int) {
	fmt.Printf("%s:\n  len=%d cap=%d x=%v\n", s, len(x), cap(x), x)
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
	printSlice("s", s)

	s = s[:0]
	printSlice("Slice the slice to give it zero length", s)

	s = s[:4]
	printSlice("Extend its length", s)

	s = s[2:]
	printSlice("Drop its first 2 values", s)

	s = s[:3]
	printSlice("Extend length to 3", s)

	// panic: runtime error: slice bounds out of range [:5] with capacity 4
	//s = s[:5]
	//printSlice("Extend length to 5", s)

	// Nil slice
	var ns []int
	printSlice("Nil slice", ns)
	//noinspection GoNilness
	if ns == nil {
		fmt.Println("s is nil")
	}

	aa := make([]int, 5)
	printSlice("aa", aa)

	bb := make([]int, 0, 5)
	printSlice("bb", bb)

	cc := bb[:2]
	printSlice("cc", cc)

	dd := cc[2:5]
	printSlice("dd", dd)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turn
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		println(strings.Join(board[i], " "))
	}

	fmt.Println("Appending to a slice:")
	var sa []int
	printSlice("Empty slice", sa)

	sa = append(sa, 0)
	printSlice("Append works on a nil slice", sa)

	sa = append(sa, 1)
	printSlice("The slice grows as needed", sa)

	sa = append(sa, 2, 3, 4)
	printSlice("We add more than one element at a time", sa)

	sa = append(sa, sa...)
	printSlice("Add slide to itself", sa)

	sa = append(sa, sa...)
	printSlice("Add slide to itself again", sa)
}

func ranges() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %3d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
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
	subtitle("Range")
	ranges()
}
