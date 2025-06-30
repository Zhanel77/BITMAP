package commands

import (
	"fmt"
	"image"
	"os"
)

// MirrorImage mirrors an image horizontally or vertically
func MirrorImage(img image.Image, mode string) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	switch mode {
	case "horizontal", "h", "horizontally", "hor":
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				p := img.At(x, y)
				dst.Set(width-x-1, y, p)
			}
		}
	case "vertical", "v", "vertically", "ver":
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				p := img.At(x, y)
				dst.Set(x, height-y-1, p)
			}
		}
	default:
		fmt.Println("Error: invalid --mirror value:", mode)
		os.Exit(1)
	}
	return dst
}

// 123
