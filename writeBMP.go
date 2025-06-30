package rwBMP

import (
	"encoding/binary"
	"image"
	"log"
	"os"
)

func WriteBMP(filename string, img image.Image) error {
	if filename == "" {
		log.Fatal("Error: No .bmp file provided\n")
	}

	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("failed to create file: %w", err)
	}
	defer outFile.Close()

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	rowSize := (width*3 + 3) &^ 3
	dataOffset := 54
	fileSize := 54 + rowSize*height

	header := make([]byte, 54)
	copy(header[:2], []byte("BM"))
	binary.LittleEndian.PutUint32(header[2:], uint32(fileSize))
	binary.LittleEndian.PutUint32(header[10:], uint32(dataOffset))
	binary.LittleEndian.PutUint32(header[14:], 40) // DIB Header size
	binary.LittleEndian.PutUint32(header[18:], uint32(width))
	binary.LittleEndian.PutUint32(header[22:], uint32(height))
	binary.LittleEndian.PutUint16(header[26:], 1)                      // number of planes
	binary.LittleEndian.PutUint16(header[28:], 24)                     // number of bytes
	binary.LittleEndian.PutUint32(header[34:], uint32(rowSize*height)) // pixel data size

	_, err = outFile.Write(header)
	if err != nil {
		log.Fatal("failed to write header: %w", err)
	}

	pixelData := make([]byte, rowSize*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, height-y-1).RGBA()
			pixelData[y*rowSize+x*3] = byte(b >> 8)
			pixelData[y*rowSize+x*3+1] = byte(g >> 8)
			pixelData[y*rowSize+x*3+2] = byte(r >> 8)
		}
	}

	_, err = outFile.Write(pixelData)
	if err != nil {
		log.Fatal("failed to write pixel data: %w", err)
	}

	return nil
}
