package filters

import (
	"image"
	"image/color"
)

func BlueFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			result.Set(x, y, color.RGBA{R: 0, G: 0, B: c.B, A: c.A})
		}
	}
	return result
}
