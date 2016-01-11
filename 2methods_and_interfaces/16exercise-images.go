package main

import "image"
import "image/color"
import "golang.org/x/tour/pic"

type Image struct {
	width  int
	height int
	color  uint8
}

// NOTE:以下の3メソッドの実装でImageインタフェースを満たしたことになる
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color, i.color, 255, 255}
}

func main() {
	m := Image{5, 5, 5}
	pic.ShowImage(m)
}
