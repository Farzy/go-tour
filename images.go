package main

import (
	"fmt"
	"image"
)

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Printf("images Bounds: %v\n", m.Bounds())
	fmt.Print("Pixel(0,0): ")
	fmt.Println(m.At(0, 0).RGBA())
}
