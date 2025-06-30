package commands

import (
	"bitmap/filters" // import your filters package (adjust path as needed)
	"errors"
	"fmt"
	"image"
)

// ApplyFilters applies all specified filters in order to the input image.
// It returns the filtered image or an error.
func ApplyFilters(img image.Image, filterList []string) (image.Image, error) {
	var err error

	for _, filterName := range filterList {
		switch filterName {
		case "blue":
			img = filters.BlueFilter(img)

		case "red":
			img = filters.RedFilter(img)

		case "green":
			img = filters.GreenFilter(img)

		case "grayscale":
			img = filters.GrayscaleFilter(img)

		case "negative":
			img = filters.NegativeFilter(img)

		case "pixelate":
			img = filters.PixelateFilter(img, 20) // default block size is 20

		case "blur":
			img = filters.BlurFilter(img)

		default:
			err = errors.New(fmt.Sprintf("unknown filter: %s", filterName))
			return nil, err
		}
		fmt.Println("Applying filter:", filterName)

	}

	return img, nil
}
