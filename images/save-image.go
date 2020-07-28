package images

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type ImageSave struct{}

func (i ImageSave) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i ImageSave) At(x, y int) color.Color {
	v1 := uint8(x ^ y)
	v2 := uint8((x + y) / 2)
	v3 := uint8(x * y)
	return color.RGBA{R: v1, G: v2, B: v3, A: 255}
}

func (i ImageSave) ColorModel() color.Model {
	return color.RGBAModel
}

func SaveImage() {
	m := ImageSave{}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, m); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	pngEncoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	f, err = os.Create("image-bestcomp.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := pngEncoder.Encode(f, m); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	f, err = os.Create("image.jpg")
	if err != nil {
		log.Fatal(err)
	}

	if err := jpeg.Encode(f, m, nil); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
