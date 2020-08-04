package allocation

import (
	"fmt"
	"go-tour/utils"
)

func Main() {

	const (
		Enone = iota + 1
		Eio
		Einval
	)

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Printf("a (%T) = %[1]v\n", a)
	fmt.Printf("s (%T) = %[1]v\n", s)
	fmt.Printf("s (%T) = %[1]v\n", m)

	const XSize, YSize = 5, 10

	utils.Subtitle("Allocate 2D slices independently")
	// Allocate the top-level slice.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, XSize)
		fmt.Printf("picture[%d]: len %d, cap %d\n", i, len(picture[i]), cap(picture[i]))
	}

	utils.Subtitle("Allocate 2D slides in one allocation")
	// Allocate the top-level slice, the same as before.
	picture2 := make([][]uint8, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture2 {
		picture2[i], pixels = pixels[:XSize], pixels[XSize:]
		fmt.Printf("picture2[%d]: len %d, cap %d\n", i, len(picture2[i]), cap(picture2[i]))
	}
}
