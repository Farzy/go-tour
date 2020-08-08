package appends

import (
	"fmt"
	"go-tour/utils"
)

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

func (p *ByteSlice) Appendp(data []byte) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
}

func Main() {
	utils.Subtitle("Append as a function")
	a := []byte{1, 2, 4, 6, 13, 27}
	fmt.Printf("a = %v, len(a) = %d, cap(a) = %d, type(a) = %[1]T\n", a, len(a), cap(a))
	a = Append(a, []byte{42})
	fmt.Printf("a = %v, len(a) = %d, cap(a) = %d, type(a) = %[1]T\n", a, len(a), cap(a))
	a = Append(a, []byte{11, 22, 33, 44, 55, 66, 77, 88})
	fmt.Printf("a = %v, len(a) = %d, cap(a) = %d, type(a) = %[1]T\n", a, len(a), cap(a))

	utils.Subtitle("Append as a method on a new type")
	b := ByteSlice{1, 2, 4, 6, 13, 27}
	fmt.Printf("b = %v, len(b) = %d, cap(b) = %d, type(b) = %[1]T\n", b, len(b), cap(b))
	b = b.Append([]byte{42})
	fmt.Printf("b = %v, len(b) = %d, cap(b) = %d, type(b) = %[1]T\n", b, len(b), cap(b))
	b = b.Append([]byte{11, 22, 33, 44, 55, 66, 77, 88})
	fmt.Printf("b = %v, len(b) = %d, cap(b) = %d, type(b) = %[1]T\n", b, len(b), cap(b))

	utils.Subtitle("Append as a method on a pointer to type")
	c := &ByteSlice{1, 2, 4, 6, 13, 27}
	fmt.Printf("c = %v, len(*c) = %d, cap(*c) = %d, type(c) = %[1]T\n", c, len(*c), cap(*c))
	c.Appendp([]byte{42})
	fmt.Printf("c = %v, len(*c) = %d, cap(*c) = %d, type(c) = %[1]T\n", c, len(*c), cap(*c))
	c.Appendp([]byte{11, 22, 33, 44, 55, 66, 77, 88})
	fmt.Printf("c = %v, len(*c) = %d, cap(*c) = %d, type(c) = %[1]T\n", c, len(*c), cap(*c))
}
