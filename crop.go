package commands

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"strconv"
)

func parseCropValue(parts []string, img *image.RGBA) (int, int, int, int) {
	var x, y, width, height int
	if len(parts) == 2 {
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
		width = img.Bounds().Dx() - x
		height = img.Bounds().Dy() - y
		if width < 0 || height < 0 {
			fmt.Println("Error: Invalid crop dimensions")
			os.Exit(1)
		}
	} else if len(parts) == 4 {
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
		width, _ = strconv.Atoi(parts[2])
		height, _ = strconv.Atoi(parts[3])
		if x+width > img.Bounds().Dx() || y+height > img.Bounds().Dy() || width < 0 || height < 0 {
			fmt.Println("Error: Invalid crop area")
			os.Exit(1)
		}
	}
	return x, y, width, height
}

func applyCropFilter(x, y, width, height int, img *image.RGBA) *image.RGBA {
	cropped := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(cropped, cropped.Bounds(), img, image.Pt(x, y), draw.Src)
	return cropped
}
