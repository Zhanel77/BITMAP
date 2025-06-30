package commands

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func Header(filename string) (error, string, uint16) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error opening file: %v", err)
	}
	// utils.CheckTrue(file)
	defer file.Close()

	headerBuffer := make([]byte, 54)

	// Read first 54 bytes
	_, err = file.Read(headerBuffer)
	if err != nil {
		log.Fatal("error reading file")
	}

	fileType := string(headerBuffer[0:2])
	fileSizeInBytes := binary.LittleEndian.Uint32(headerBuffer[2:6])
	headerSize := uint32(14)
	dibHeaderSize := binary.LittleEndian.Uint32(headerBuffer[14:18])
	widthInPixels := int32(binary.LittleEndian.Uint32(headerBuffer[18:22]))
	heightInPixels := int32(binary.LittleEndian.Uint32(headerBuffer[22:26]))
	pixelSizeInBits := binary.LittleEndian.Uint16(headerBuffer[28:30])
	imageSizeInBytes := binary.LittleEndian.Uint32(headerBuffer[34:38])

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("не удалось получить информацию о файле: %v", err)
	}
	actualFileSize := uint32(fileInfo.Size())

	if fileSizeInBytes != actualFileSize {
		log.Fatal("фактический размер файла (%d байт) не соответствует заявленному в заголовке (%d байт)",
			actualFileSize, fileSizeInBytes)
	}

	// Формируем строку для вывода информации
	headerString := fmt.Sprintf(`+---------------------------+
|     Bitmap File Header    |
+---------------------------+

BMP Header:
- FileType: %s
- FileSizeInBytes: %d
- HeaderSize: %d
DIB Header:
- DibHeaderSize: %d
- WidthInPixels: %d
- HeightInPixels: %d
- PixelSizeInBits: %d
- ImageSizeInBytes: %d
`,
		fileType,
		fileSizeInBytes,
		headerSize,
		dibHeaderSize,
		widthInPixels,
		heightInPixels,
		pixelSizeInBits,
		imageSizeInBytes,
	)

	return nil, headerString, pixelSizeInBits
}
