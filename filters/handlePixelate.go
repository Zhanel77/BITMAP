// handlePixelate.go
package filters

import (
	"image"
	"image/color"
)

func PixelateFilter(img image.Image, blockSize int) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
			var r, g, b, count int
			for dy := 0; dy < blockSize && y+dy < bounds.Max.Y; dy++ {
				for dx := 0; dx < blockSize && x+dx < bounds.Max.X; dx++ {
					c := color.RGBAModel.Convert(img.At(x+dx, y+dy)).(color.RGBA)
					r += int(c.R)
					g += int(c.G)
					b += int(c.B)
					count++
				}
			}
			if count == 0 {
				continue
			}
			avg := color.RGBA{R: uint8(r / count), G: uint8(g / count), B: uint8(b / count), A: 255}
			for dy := 0; dy < blockSize && y+dy < bounds.Max.Y; dy++ {
				for dx := 0; dx < blockSize && x+dx < bounds.Max.X; dx++ {
					result.Set(x+dx, y+dy, avg)
				}
			}
		}
	}
	return result
}
