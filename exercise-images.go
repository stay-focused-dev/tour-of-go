package main

import (
	"golang.org/x/tour/pic"

	"image"
	"image/color"
)

type MyImage struct {
	X, Y int
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.X, i.Y)
}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ 256), uint8(y ^ 256), 255, 255}
}

func main() {
	m := MyImage{256, 256}
	pic.ShowImage(m)
}
