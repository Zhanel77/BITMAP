package main

import (
	"bitmap/commands"
	"bitmap/utils"
	"fmt"
	"log"
	"os"
	"strings"

	. "bitmap/rwBMP"
)

func main() {
	args := os.Args

	if len(args) < 2 || args[1] == "--help" || args[1] == "-h" {
		utils.PrintGeneralHelp()
		return
	}

	command := args[1]

	for _, arg := range args[2:] {
		if arg == "--helps" || arg == "-h" {
			switch command {
			case "header":
				utils.PrintHeaderHelp()
			case "apply":
				utils.PrintApplyHelp()
			default:
				utils.PrintGeneralHelp()

			}
			return
		}
	}

	// utils.Help(args)

	var filename string
	for i := 2; i < len(args); i++ {
		if strings.HasSuffix(args[i], ".bmp") {
			filename = args[i]
			break
		}
	}

	if filename == "" {
		log.Fatal("Error: No .bmp file provided\n")
	}
	if err := ValidateBMP(filename); err != nil {
		log.Fatal(err)
	}

	switch command {
	case "header":
		if err, headerRes, _ := commands.Header(filename); err != nil {
			log.Fatalf("Error: %v\n", err)
		} else {
			fmt.Println(headerRes)
		}

	case "apply":
		commands.HandleApply(os.Args[2:])
	default:
		log.Fatalf("Invalid command:%s\n", command)
	}
}
