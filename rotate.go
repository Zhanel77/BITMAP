package commands

import (
	"image"
)

func RotateImage(img image.Image, angle int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	var rotated *image.RGBA

	switch angle {
	case 90:
		rotated = image.NewRGBA(image.Rect(0, 0, height, width))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rotated.Set(height-y-1, x, img.At(x, y))
			}
		}
	case 180:
		rotated = image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rotated.Set(width-x-1, height-y-1, img.At(x, y))
			}
		}
	case 270:
		rotated = image.NewRGBA(image.Rect(0, 0, height, width))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rotated.Set(y, width-x-1, img.At(x, y))
			}
		}
	default:
		return img
	}
	return rotated
}
