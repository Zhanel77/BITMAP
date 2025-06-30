package filters

import (
	"image"
	"image/color"
)

func NegativeFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			result.Set(x, y, color.RGBA{R: 255 - c.R, G: 255 - c.G, B: 255 - c.B, A: c.A})
		}
	}
	return result
}
