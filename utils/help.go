package utils

import "fmt"

func PrintGeneralHelp() {
	fmt.Println(`Usage:
  bitmap <command> [arguments]

The commands are:
  header    prints bitmap file header information
  apply     applies processing to the image and saves it to the file`)
}

func PrintHeaderHelp() {
	fmt.Println(`Usage:
  bitmap header <source_file>

Description:
  Prints bitmap file header information`)
}

func PrintApplyHelp() {
	fmt.Println(`Usage:
  bitmap apply [options] <source_file> <output_file>

The options are:
  -h, --help      prints program usage information
  -mh             mirror horizontally
  -mv             mirror vertically
  -rr             rotate right
  -rl             rotate left
  -fg             grayscale filter
  -fn             negative filter`)
}
