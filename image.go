package main

import (
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) At(xm y int) color.Color {
	c := uint8(2*(x-x*x)-2*(y*y-y))
	return color.RGBA(c, c, 255, 255)
}

func main() {
	m := Image{480, 360}
	pic.Show(m)
}
