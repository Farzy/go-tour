package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			val := x * y
			//val := (x^y) * 4
			//val := x^y
			//val := (x + y) / 2
			//fmt.Printf("(%v,%v): val = %v, val uint8 = %v\n", x, y, val, uint8(val))
			picture[y][x] = uint8(val)
		}
	}

	return picture
}

func exerciceSlices() {
	pic.Show(Pic)
}
