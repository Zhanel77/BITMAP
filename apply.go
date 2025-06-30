package commands

import (
	"bitmap/rwBMP"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func HandleApply(args []string) {
	var rotateAngles []int
	var srcFile, outFile string
	var cropValues [][]string
	var mirrorModes []string
	var filtersToApply []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--rotate=") {
			value := strings.TrimPrefix(arg, "--rotate=")
			switch value {
			case "right", "90":
				rotateAngles = append(rotateAngles, 90)
			case "180", "-180":
				rotateAngles = append(rotateAngles, 180)
			case "left", "-90", "270":
				rotateAngles = append(rotateAngles, 270)
			case "-270":
				rotateAngles = append(rotateAngles, 90)
			default:
				angle, err := strconv.Atoi(value)
				if err != nil || (angle != 90 && angle != 180 && angle != 270 && angle != -90 && angle != -180 && angle != -270) {
					fmt.Println("Error: invalid --rotate value")
					os.Exit(1)
				}
				if angle < 0 {
					angle = 360 + angle
				}
				rotateAngles = append(rotateAngles, angle)
			}
		} else if strings.HasPrefix(arg, "--crop=") {
			value := strings.TrimPrefix(arg, "--crop=")
			cropParts := strings.Split(value, "-")
			if len(cropParts) != 2 && len(cropParts) != 4 {
				fmt.Println("Error: invalid --crop value, must be 2 or 4 numbers")
				os.Exit(1)
			}
			cropValues = append(cropValues, cropParts)
		} else if strings.HasPrefix(arg, "--mirror=") {
			value := strings.TrimPrefix(arg, "--mirror=")
			mirrorModes = append(mirrorModes, strings.ToLower(value))
		} else if strings.HasPrefix(arg, "--filter=") {
			value := strings.TrimPrefix(arg, "--filter=")
			filtersToApply = append(filtersToApply, strings.ToLower(value))

		} else if strings.HasSuffix(arg, ".bmp") {
			if srcFile == "" {
				srcFile = arg
			} else {
				outFile = arg
			}
		}
	}

	if srcFile == "" || outFile == "" {
		fmt.Println("Error: missing input/output file")
		os.Exit(1)
	}

	img, err := rwBMP.ReadBMP(srcFile)
	if err != nil {
		fmt.Println("Error reading BMP:", err)
		os.Exit(1)
	}

	imgRGBA, ok := img.(*image.RGBA)
	if !ok {
		fmt.Println("Error: image is not in RGBA format")
		os.Exit(1)
	}

	for _, parts := range cropValues {
		x, y, width, height := parseCropValue(parts, imgRGBA)
		imgRGBA = applyCropFilter(x, y, width, height, imgRGBA)
	}

	for _, mode := range mirrorModes {
		imgRGBA = MirrorImage(imgRGBA, mode).(*image.RGBA)
	}

	for _, angle := range rotateAngles {
		imgRGBA = RotateImage(imgRGBA, angle).(*image.RGBA)
	}
	if len(filtersToApply) > 0 {
		imgFiltered, err := ApplyFilters(imgRGBA, filtersToApply)
		if err != nil {
			fmt.Println("Error applying filters:", err)
			os.Exit(1)
		}
		// Convert to *image.RGBA if needed
		if rgba, ok := imgFiltered.(*image.RGBA); ok {
			imgRGBA = rgba
		} else {
			// Create RGBA from generic image.Image
			bounds := imgFiltered.Bounds()
			tmp := image.NewRGBA(bounds)
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					tmp.Set(x, y, imgFiltered.At(x, y))
				}
			}
			imgRGBA = tmp
		}
	}

	err = rwBMP.WriteBMP(outFile, imgRGBA)
	if err != nil {
		fmt.Println("Error writing BMP:", err)
		os.Exit(1)
	}

	fmt.Println("Image processed and saved:", outFile)
}
