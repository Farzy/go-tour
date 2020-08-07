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

func Main() {
	utils.Subtitle("Testing different versions of Append")

	a := []byte{1, 2, 4, 6, 13, 27}
	fmt.Printf("a = %v, len(a) = %d, cap(a) = %d, type(a) = %[1]T\n", a, len(a), cap(a))
	a = Append(a, []byte{42})
	fmt.Printf("a = %v, len(a) = %d, cap(a) = %d, type(a) = %[1]T\n", a, len(a), cap(a))
}
