package rwBMP

import (
	"encoding/binary"
	"image"
	"image/color"
	"log"
	"os"
)

func ReadBMP(filename string) (image.Image, error) {
	err := ValidateBMP(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Не удалось открыть файл: %w", err)
	}
	defer file.Close()

	var header [54]byte
	_, err = file.Read(header[:])
	if err != nil {
		log.Fatal("не удалось прочитать заголовок BMP: %w", err)
	}

	dataOffset := binary.LittleEndian.Uint32(header[10:14])
	width := int(binary.LittleEndian.Uint32(header[18:22]))
	height := int(binary.LittleEndian.Uint32(header[22:26]))
	if height < 0 {
		height = -height
	}

	file.Seek(int64(dataOffset), 0)

	rowSize := (width*3 + 3) &^ 3
	pixelData := make([]byte, rowSize*height)
	_, err = file.Read(pixelData)
	if err != nil {
		log.Fatal("Не удалось прочитать пиксели файла: %w", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := (y*rowSize + x*3)
			if i+2 >= len(pixelData) {
				log.Fatal("Индекс пикселя находится за пределами границ данных")
			}
			b := pixelData[i]
			g := pixelData[i+1]
			r := pixelData[i+2]
			img.Set(x, height-y-1, color.RGBA{R: r, G: g, B: b, A: 255})

		}
	}
	return img, nil
}
