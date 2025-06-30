// handleBlur.go
package filters

import (
	"image"
	"image/color"
)

func BlurFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b, count int
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					xx := x + dx
					yy := y + dy
					if xx >= bounds.Min.X && xx < bounds.Max.X && yy >= bounds.Min.Y && yy < bounds.Max.Y {
						c := color.RGBAModel.Convert(img.At(xx, yy)).(color.RGBA)
						r += int(c.R)
						g += int(c.G)
						b += int(c.B)
						count++
					}
				}
			}
			result.Set(x, y, color.RGBA{R: uint8(r / count), G: uint8(g / count), B: uint8(b / count), A: 255})
		}
	}
	return result
}
