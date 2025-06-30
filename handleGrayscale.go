// handleGrayscale.go
package filters

import (
	"image"
	"image/color"
)

func GrayscaleFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			gray := uint8(0.299*float32(c.R) + 0.587*float32(c.G) + 0.114*float32(c.B))
			result.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: c.A})
		}
	}
	return result
}
