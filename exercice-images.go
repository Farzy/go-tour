package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
	v1 := uint8(x ^ y)
	v2 := uint8((x + y) / 2)
	v3 := uint8(x * y)
	return color.RGBA{R: v1, G: v2, B: v3, A: 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func exerciceImages() {
	m := Image{}
	pic.ShowImage(m)
}
