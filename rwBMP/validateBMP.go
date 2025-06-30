package rwBMP

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ValidateBMP(filePath string) error {
	if ext := strings.ToLower(filepath.Ext(filePath)); ext != ".bmp" {
		log.Fatal("%s is not a bitmap file\n", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer file.Close()

	// Buffer to read file type
	headerBuffer := make([]byte, 2)

	// Read first 2 bytes
	_, err = file.Read(headerBuffer)
	if err != nil {
		log.Fatal("error reading file: %v", err)
	}

	// Check BMP signature
	fileType := string(headerBuffer)
	if fileType != "BM" {
		log.Fatal("is not a bitmap file", filePath)
	}

	return nil
}
